package repository

import (
	"context"
	"errors"

	"github.com/example/hungry-hub/internal/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MenuItemRepo struct {
	db *gorm.DB
}

func NewMenuItemRepo(db *gorm.DB) *MenuItemRepo {
	return &MenuItemRepo{db: db}
}

func (m *MenuItemRepo) Create(ctx context.Context, restaurantID int64, name string, description *string, price decimal.Decimal, category *string, isAvailable *bool) (models.MenuItem, error) {
	avail := true
	if isAvailable != nil {
		avail = *isAvailable
	}

	item := models.MenuItem{
		RestaurantID: restaurantID,
		Name:         name,
		Description:  description,
		Price:        price,
		Category:     category,
		IsAvailable:  avail,
	}

	if err := m.db.WithContext(ctx).Create(&item).Error; err != nil {
		return models.MenuItem{}, err
	}

	return m.Get(ctx, item.ID)
}

func (m *MenuItemRepo) ListByRestaurant(ctx context.Context, restaurantID int64, category *string) ([]models.MenuItem, error) {
	var out []models.MenuItem
	q := m.db.WithContext(ctx).Where("restaurant_id = ?", restaurantID)
	if category != nil && *category != "" {
		q = q.Where("category = ?", *category)
	}

	err := q.Order("id ASC").Find(&out).Error

	return out, err
}

func (m *MenuItemRepo) Get(ctx context.Context, id int64) (models.MenuItem, error) {
	var item models.MenuItem
	err := m.db.WithContext(ctx).First(&item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.MenuItem{}, ErrNotFound
	}

	if err != nil {
		return models.MenuItem{}, err
	}

	return item, nil
}

func (m *MenuItemRepo) Update(ctx context.Context, id int64, name, description *string, price *decimal.Decimal, category **string, isAvailable *bool) (models.MenuItem, error) {
	updates := map[string]any{}
	if name != nil {
		updates["name"] = *name
	}

	if description != nil {
		updates["description"] = *description
	}

	if price != nil {
		updates["price"] = *price
	}

	if category != nil {
		updates["category"] = derefOrNil(category)
	}

	if isAvailable != nil {
		updates["is_available"] = *isAvailable
	}

	if len(updates) == 0 {
		return m.Get(ctx, id)
	}

	res := m.db.WithContext(ctx).Model(&models.MenuItem{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return models.MenuItem{}, res.Error
	}

	if res.RowsAffected == 0 {
		return models.MenuItem{}, ErrNotFound
	}

	return m.Get(ctx, id)
}

func (m *MenuItemRepo) Delete(ctx context.Context, id int64) error {
	res := m.db.WithContext(ctx).Delete(&models.MenuItem{}, id)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

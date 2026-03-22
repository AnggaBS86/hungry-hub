package repository

import (
	"context"
	"errors"

	"github.com/example/hungry-hub/internal/models"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type RestaurantRepo struct {
	db *gorm.DB
}

func NewRestaurantRepo(db *gorm.DB) *RestaurantRepo {
	return &RestaurantRepo{db: db}
}

func (r *RestaurantRepo) Create(ctx context.Context, name, address string, phone, openingHours *string) (models.Restaurant, error) {
	item := models.Restaurant{
		Name:         name,
		Address:      address,
		Phone:        phone,
		OpeningHours: openingHours,
	}

	if err := r.db.WithContext(ctx).Create(&item).Error; err != nil {
		return models.Restaurant{}, err
	}

	return r.Get(ctx, item.ID)
}

func (r *RestaurantRepo) List(ctx context.Context) ([]models.Restaurant, error) {
	var out []models.Restaurant
	err := r.db.WithContext(ctx).
		Order("id ASC").
		Find(&out).Error

	return out, err
}

func (r *RestaurantRepo) Get(ctx context.Context, id int64) (models.Restaurant, error) {
	var item models.Restaurant
	err := r.db.WithContext(ctx).First(&item, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Restaurant{}, ErrNotFound
	}

	if err != nil {
		return models.Restaurant{}, err
	}

	return item, nil
}

func (r *RestaurantRepo) Update(ctx context.Context, id int64, name, address *string, phone, openingHours **string) (models.Restaurant, error) {
	updates := map[string]any{}
	if name != nil {
		updates["name"] = *name
	}

	if address != nil {
		updates["address"] = *address
	}

	if phone != nil {
		updates["phone"] = derefOrNil(phone)
	}

	if openingHours != nil {
		updates["opening_hours"] = derefOrNil(openingHours)
	}

	if len(updates) == 0 {
		return r.Get(ctx, id)
	}

	res := r.db.WithContext(ctx).Model(&models.Restaurant{}).Where("id = ?", id).Updates(updates)
	if res.Error != nil {
		return models.Restaurant{}, res.Error
	}

	if res.RowsAffected == 0 {
		return models.Restaurant{}, ErrNotFound
	}

	return r.Get(ctx, id)
}

func (r *RestaurantRepo) Delete(ctx context.Context, id int64) error {
	res := r.db.WithContext(ctx).Delete(&models.Restaurant{}, id)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func derefOrNil(ptr any) any {
	switch v := ptr.(type) {
	case **string:
		if v == nil || *v == nil {
			return nil
		}
		return **v
	case *string:
		if v == nil {
			return nil
		}
		return *v
	default:
		return nil
	}
}

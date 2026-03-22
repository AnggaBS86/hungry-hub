package seed

import (
	"context"
	"fmt"

	"github.com/example/hungry-hub/internal/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	var count int64
	if err := db.WithContext(context.Background()).Model(&models.Restaurant{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	ctx := context.Background()
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		type rest struct {
			name    string
			address string
			phone   *string
			hours   *string
		}

		r1Phone := "+66-2-111-1111"
		r1Hours := "Mon-Sun 10:00-22:00"
		r2Phone := "+66-2-222-2222"
		r2Hours := "Mon-Fri 09:00-18:00"
		restaurants := []rest{
			{name: "Bangkok Bistro", address: "123 Sukhumvit Rd, Bangkok", phone: &r1Phone, hours: &r1Hours},
			{name: "Chiang Mai Cafe", address: "45 Nimman Rd, Chiang Mai", phone: &r2Phone, hours: &r2Hours},
		}

		var ids []int64
		for _, r := range restaurants {
			item := models.Restaurant{
				Name:         r.name,
				Address:      r.address,
				Phone:        r.phone,
				OpeningHours: r.hours,
			}

			if err := tx.Create(&item).Error; err != nil {
				return err
			}

			ids = append(ids, item.ID)
		}

		items := []struct {
			restaurantIndex int
			name            string
			desc            string
			price           string
			category        string
			available       bool
		}{
			{0, "Spring Rolls", "Crispy veggie spring rolls", "120.00", "appetizer", true},
			{0, "Pad Thai", "Classic stir-fried noodles", "180.00", "main", true},
			{0, "Green Curry", "Spicy coconut curry", "220.00", "main", true},
			{0, "Mango Sticky Rice", "Sweet mango dessert", "140.00", "dessert", true},
			{0, "Thai Iced Tea", "Sweet milk tea", "80.00", "drink", true},
			{1, "Caesar Salad", "Romaine, parmesan, croutons", "150.00", "appetizer", true},
			{1, "Grilled Chicken", "Served with seasonal vegetables", "250.00", "main", true},
			{1, "Tom Yum Soup", "Hot and sour soup", "160.00", "main", true},
			{1, "Chocolate Cake", "Rich cocoa cake", "130.00", "dessert", true},
			{1, "Espresso", "Single shot espresso", "70.00", "drink", true},
		}

		for _, item := range items {
			p, _ := decimal.NewFromString(item.price)
			restaurantID := ids[item.restaurantIndex]
			desc := item.desc
			category := item.category
			menuItem := models.MenuItem{
				RestaurantID: restaurantID,
				Name:         item.name,
				Description:  &desc,
				Price:        p,
				Category:     &category,
				IsAvailable:  item.available,
			}

			if err := tx.Create(&menuItem).Error; err != nil {
				return err
			}
		}

		fmt.Println("seed completed")

		return nil
	})
}

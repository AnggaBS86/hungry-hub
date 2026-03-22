package models

import "time"

type Restaurant struct {
	ID           int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string     `json:"name" gorm:"type:varchar(255);not null"`
	Address      string     `json:"address" gorm:"type:varchar(255);not null"`
	Phone        *string    `json:"phone,omitempty" gorm:"type:varchar(50)"`
	OpeningHours *string    `json:"opening_hours,omitempty" gorm:"type:varchar(255)"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	MenuItems    []MenuItem `json:"menu_items,omitempty" gorm:"foreignKey:RestaurantID"`
}

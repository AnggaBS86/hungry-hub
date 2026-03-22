package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type MenuItem struct {
	ID           int64           `json:"id" gorm:"primaryKey;autoIncrement"`
	RestaurantID int64           `json:"restaurant_id" gorm:"not null;index"`
	Name         string          `json:"name" gorm:"type:varchar(255);not null;index"`
	Description  *string         `json:"description,omitempty" gorm:"type:text"`
	Price        decimal.Decimal `json:"price" gorm:"type:decimal(10,2);not null"`
	Category     *string         `json:"category,omitempty" gorm:"type:varchar(50);index"`
	IsAvailable  bool            `json:"is_available" gorm:"not null;default:true"`
	CreatedAt    time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
}

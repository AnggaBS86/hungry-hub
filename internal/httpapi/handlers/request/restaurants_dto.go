package request

type CreateRestaurantDTO struct {
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	Phone        *string `json:"phone"`
	OpeningHours *string `json:"opening_hours"`
}

type UpdateRestaurantDTO struct {
	Name         *string  `json:"name"`
	Address      *string  `json:"address"`
	Phone        **string `json:"phone"`
	OpeningHours **string `json:"opening_hours"`
}


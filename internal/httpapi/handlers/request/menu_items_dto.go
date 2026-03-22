package request

type CreateMenuItemDTO struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       string  `json:"price"`
	Category    *string `json:"category"`
	IsAvailable *bool   `json:"is_available"`
}

type UpdateMenuItemDTO struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *string  `json:"price"`
	Category    **string `json:"category"`
	IsAvailable *bool    `json:"is_available"`
}


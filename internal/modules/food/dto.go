package food

type CreateFoodRequest struct {
	MenuID int64   `json:"menu_id"`
	Name   string  `json:"name" validate:"required,min=2,max=100"`
	Price  float64 `json:"price" validate:"required,gte=0"`
	Image  *string `json:"image,omitempty"`
}

type UpdateFoodRequest struct {
	MenuID int64   `json:"menu_id"`
	Name   string  `json:"name" validate:"required,min=2,max=100"`
	Price  float64 `json:"price" validate:"required,gte=0"`
	Image  *string `json:"image,omitempty"`
}

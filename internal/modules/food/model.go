package food

import "time"

type Food struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" validate:"required,min=2,max=100"`
	Price     float64   `json:"price" validate:"required"`
	Image     *string   `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	MenuID    int64     `json:"menu_id" validate:"required"`
}

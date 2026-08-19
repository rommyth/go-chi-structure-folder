package order

import (
	"restaurant-management/internal/modules/table"
	"time"
)

type Order struct {
	ID        int64        `json:"id"`
	TableID   int64        `json:"table_id" validate:"required"`
	Notes     *string      `json:"notes"`
	OrderDate time.Time    `json:"order_date" validate:"required"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Table     *table.Table `json:"table,omitempty"`
}

type OrderItem struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id" validate:"required"`
	FoodID    int64     `json:"food_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,gt=0"`
	UnitPrice float64   `json:"unit_price" validate:"required,gte=0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

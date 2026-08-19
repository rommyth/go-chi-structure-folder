package invoice

import "time"

type Invoice struct {
	ID             int64     `json:"id"`
	OrderID        int64     `json:"order_id"`
	PaymentMethod  string    `json:"payment_method" validate:"oneof=CASH CREDIT"`
	PaymentStatus  string    `json:"payment_status" validate:"oneof=CANCELED PENDING PAID"`
	PaymentDueDate time.Time `json:"payment_due_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

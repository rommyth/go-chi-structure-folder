package table

import "time"

type Table struct {
	ID             int64     `json:"id"`
	NumberOfGuests int       `json:"number_of_guests" validate:"required"`
	TableNumber    int       `json:"table_number" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

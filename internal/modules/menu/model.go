package menu

import "time"

type Menu struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name" validate:"required,min=2,max=100"`
	Category  string     `json:"category"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

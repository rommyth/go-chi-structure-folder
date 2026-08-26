package user

import "time"

type User struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name" validate:"required,min=2,max=100"`
	LastName     string    `json:"last_name" validate:"required,min=2,max=100"`
	Password     string    `json:"password,omitempty" validate:"required,min=6"`
	Email        string    `json:"email" validate:"required,email"`
	Avatar       *string   `json:"avatar"`
	Phone        string    `json:"phone" validate:"required"`
	Token        string    `json:"-"`
	RefreshToken string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

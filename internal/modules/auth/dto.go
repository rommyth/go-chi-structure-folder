package auth

import "restaurant-management/internal/modules/user"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	FirstName string  `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string  `json:"last_name" validate:"required,min=2,max=100"`
	Email     string  `json:"email" validate:"required,email"`
	Password  string  `json:"password" validate:"required,min=6"`
	Phone     string  `json:"phone" validate:"required"`
	Avatar    *string `json:"avatar"`
}

type LoginResponse struct {
	Token        string     `json:"token"`
	RefreshToken string     `json:"refresh_token"`
	User         *user.User `json:"user,omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

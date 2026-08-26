package user

type UpdateUserRequest struct {
	FirstName string  `json:"first_name" validate:"required,min=3"`
	LastName  string  `json:"last_name" validate:"required,min=3"`
	Email     string  `json:"email" validate:"required,min=3,email"`
	Phone     string  `json:"phone" validate:"required"`
	Avatar    *string `json:"avatar"`
}

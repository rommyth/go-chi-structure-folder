package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type CustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

func GetClaims(ctx context.Context) (*CustomClaims, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("invalid user_id claim")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid email claim")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid role claim")
	}

	return &CustomClaims{
		UserID: int64(userID),
		Role:   role,
		Email:  email,
	}, nil
}

func (c *CustomClaims) GenerateToken(jwt *jwtauth.JWTAuth, exp time.Time) (t jwt.Token, tokenString string, err error) {
	return jwt.Encode(map[string]interface{}{
		"user_id": c.UserID,
		"role":    c.Role,
		"email":   c.Email,
		"exp":     exp.Unix(),
	})
}

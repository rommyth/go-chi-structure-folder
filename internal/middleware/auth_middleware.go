package middleware

import (
	"errors"
	"net/http"
	"restaurant-management/internal/modules/auth"
	"restaurant-management/pkg/response"

	"github.com/go-chi/jwtauth/v5"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
)

func Verify(jwt *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return jwtauth.Verifier(jwt)
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil || token == nil {
			response.Error(
				w,
				http.StatusUnauthorized,
				ErrUnauthorized.Error(),
				nil,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetClaims(r *http.Request) (*auth.CustomClaims, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return nil, err
	}

	userID, ok := claims["user_id"].(string)
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

	return &auth.CustomClaims{
		UserID: userID,
		Role:   role,
		Email:  email,
	}, nil
}

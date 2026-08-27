package middleware

import (
	"errors"
	"net/http"
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

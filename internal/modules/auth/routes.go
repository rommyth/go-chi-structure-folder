package auth

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/signup", h.SignUp)
		r.Post("/refresh-token", h.RefreshToken)
	})
}

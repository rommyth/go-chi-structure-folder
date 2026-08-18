package user

import (
	"net/http"
	"restaurant-management/pkg/response"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			response.Success(w, http.StatusOK, "Success Test Get User", map[string]string{
				"user_id": "12346",
				"name":    "Romi",
			})
		})
	})
}

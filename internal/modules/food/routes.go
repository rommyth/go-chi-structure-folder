package food

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/foods", func(r chi.Router) {
		r.Get("/", h.GetFoods)
		r.Get("/{id}", h.GetFoodByID)
		r.Post("/", h.CreateFood)
		r.Patch("/{id}", h.UpdateFood)
	})
}

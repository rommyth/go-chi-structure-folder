package table

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/orders", func(r chi.Router) {
		r.Get("/", h.GetTables)
		r.Get("/{id}", h.GetTableByID)
		r.Post("/", h.CreateTable)
		r.Put("/{id}", h.UpdateTable)
	})
}

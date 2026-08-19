package menu

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/menus", func(r chi.Router) {
		r.Get("/", h.GetMenus)
		r.Get("/{id}", h.GetMenuByID)
		r.Post("/", h.CreateMenu)
		r.Put("/{id}", h.UpdateMenu)
	})
}

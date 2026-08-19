package invoice

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/invoices", func(r chi.Router) {
		r.Get("/", h.GetInvoices)
		r.Get("/{id}", h.GetInvoiceByID)
		r.Post("/", h.CreateInvoice)
		r.Put("/{id}", h.UpdateInvoice)
	})
}

package order

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/orders", func(r chi.Router) {
		r.Get("/", h.GetOrders)
		r.Get("/{id}", h.GetOrderByID)
		r.Post("/", h.CreateOrder)
		r.Put("/{id}", h.UpdateOrder)
	})

	r.Route("/order-items", func(r chi.Router) {
		r.Get("/", h.GetOrderItems)
		r.Get("/{id}", h.GetOrderItemByID)
		r.Get("/order/{order_id}", h.GetOrderItemsByOrderID)
		r.Post("/", h.CreateOrderItems)
		r.Put("/{id}", h.UpdateOrderItem)
	})
}

package order

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetOrderItems(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetOrderItemByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetOrderItemsByOrderID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) CreateOrderItems(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	// TODO
}

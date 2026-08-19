package invoice

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetInvoices(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetInvoiceByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	// TODO
}

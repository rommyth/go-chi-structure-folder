package table

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTables(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetTableByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) CreateTable(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) UpdateTable(w http.ResponseWriter, r *http.Request) {
	// TODO
}

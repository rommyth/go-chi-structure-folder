package menu

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMenus(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetMenuByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	// TODO
}

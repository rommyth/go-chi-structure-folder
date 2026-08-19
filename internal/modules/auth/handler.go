package auth

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO
}

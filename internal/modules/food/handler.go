package food

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetFoods(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetFoodByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) CreateFood(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) UpdateFood(w http.ResponseWriter, r *http.Request) {
	// TODO
}

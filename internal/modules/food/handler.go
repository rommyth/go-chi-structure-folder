package food

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"restaurant-management/pkg/response"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  *Service
	log      *slog.Logger
	validate *validator.Validate
}

func NewHandler(service *Service, log *slog.Logger, validate *validator.Validate) *Handler {
	return &Handler{
		service:  service,
		log:      log,
		validate: validate,
	}
}

func (h *Handler) GetFoods(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetFoodByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) CreateFood(w http.ResponseWriter, r *http.Request) {
	var req CreateFoodRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.WarnContext(r.Context(), "invalid parsing request", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		h.log.WarnContext(r.Context(), "invalid validate", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	result, err := h.service.Create(r.Context(), req)
	if err != nil {
		h.log.WarnContext(r.Context(), "failed create food", "error", err)
		response.Error(w, http.StatusInternalServerError, "failed to create food", err)
		return
	}

	h.log.InfoContext(r.Context(), "food created", "food_id", result.ID)
	response.Success(w, http.StatusCreated, "success create food", result)

}

func (h *Handler) UpdateFood(w http.ResponseWriter, r *http.Request) {
	// TODO
}

package food

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"restaurant-management/pkg/response"
	"strconv"

	"github.com/go-chi/chi/v5"
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

// GetFoods godoc
// @Summary 	Get all foods
// @Produce 	json
// @Tags 		Foods
// @Param		page query int false "Page Number" default(1)
// @Param		limit query int false "Limit NUmber" default(10)
// @Param		search query string false "Search"
// @Success		200 (object) response.Success
// @Router		/foods [get]
func (h *Handler) GetFoods(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 10

	if value := r.URL.Query().Get("page"); value != "" {
		var err error
		page, err = strconv.Atoi(value)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid page", err)
		}
	}

	if value := r.URL.Query().Get("limit"); value != "" {
		var err error
		page, err = strconv.Atoi(value)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid limit", err)
		}
	}
	search := r.URL.Query().Get("search")

	foods, total, err := h.service.GetList(r.Context(), page, limit, search)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to get list food", err)
		return
	}

	response.SuccessWithMeta(w, http.StatusOK, "Success get foods", foods, page, limit, total)
}

// GetFoodByID godoc
// @Summary 	Get single food by id
// @Produce 	json
// @Tags 		Foods
// @Param		id path int true "food_id"
// @Success		200 (object) response.Success
// @Router		/foods/{id} [get]
func (h *Handler) GetFoodByID(w http.ResponseWriter, r *http.Request) {
	foodId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid food id", err)
		return
	}

	food, err := h.service.GetByID(r.Context(), foodId)
	if err != nil {
		if errors.Is(err, ErrFoodNotFound) {
			response.Error(w, http.StatusNotFound, "food not found", err)
			return
		}

		response.Error(w, http.StatusInternalServerError, "failed get food", err)
		return
	}

	response.Success(w, http.StatusOK, "success get food", food)
}

// CreateFood godoc
// @Summary 	Create Food
// @Produce 	json
// @Tags 		Foods
// @Param		request body CreateFoodRequest true "food data"
// @Success		200 (object) response.Success
// @Router		/foods [post]
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

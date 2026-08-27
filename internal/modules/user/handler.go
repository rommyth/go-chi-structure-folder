package user

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

// GetUsers godoc
// @Summary 	List User
// @Produce		json
// @Tags		Users
// @Param		page query int false "page number" default(1)
// @Param		limit query int false "limit number" default(10)
// @Param		search query string false "search user by name"
// @Success		200 {object} response.Response
// @Router		/users [get]
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 10

	if value := r.URL.Query().Get("page"); value != "" {
		var err error
		page, err = strconv.Atoi(value)
		if err != nil {
			h.log.WarnContext(r.Context(), "page required", "error", err)
			response.Error(w, http.StatusBadRequest, "page required", err)
			return
		}
	}

	if value := r.URL.Query().Get("limit"); value != "" {
		var err error
		page, err = strconv.Atoi(value)
		if err != nil {
			h.log.WarnContext(r.Context(), "limit required", "error", err)
			response.Error(w, http.StatusBadRequest, "limit required", err)
			return
		}
	}

	search := r.URL.Query().Get("search")

	users, total, err := h.service.GetUsers(r.Context(), page, limit, search)
	if err != nil {
		h.log.WarnContext(r.Context(), "failed get users", "error", err)
		response.Error(w, http.StatusInternalServerError, "failed get users", err)
		return
	}

	h.log.InfoContext(r.Context(), "success get users")
	response.SuccessWithMeta(w, http.StatusInternalServerError, "success get users", users, page, limit, total)
}

// GetUserByID 	godoc
// @Summary 	Get User By ID
// @Produce		json
// @Tags		Users
// @Param		id path int true "user id"
// @Success		200 {object} response.Response(meta=null)
// @Router		/users/{id} [get]
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.log.WarnContext(r.Context(), "invalid id", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid id", err)
		return
	}

	user, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			h.log.WarnContext(r.Context(), "user not found", "error", err)
			response.Error(w, http.StatusNotFound, "user not found", err)
			return
		}

		h.log.WarnContext(r.Context(), "failed get user", "error", err)
		response.Error(w, http.StatusNotFound, "failed get user", err)
		return
	}

	h.log.InfoContext(r.Context(), "success get user", "user_id", id)
	response.Success(w, http.StatusOK, "success get user", user)
}

// UpdateUser 	godoc
// @Summary 	Update User
// @Produce		json
// @Tags		Users
// @Param		id path int true "user id"
// @Param		requrest body UpdateUserRequest true "user body update"
// @Success		200 {object} response.Response
// @Router		/users/{id} [patch]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.log.WarnContext(r.Context(), "invalid id", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid id", err)
		return
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.WarnContext(r.Context(), "failed decode request", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid validation", err)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		h.log.WarnContext(r.Context(), "invalid validation", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid validation", err)
		return
	}

	user, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			h.log.WarnContext(r.Context(), "user not found", "error", err)
			response.Error(w, http.StatusNotFound, "user not found", err)
			return
		}

		if errors.Is(err, ErrEmailAlreadyExist) {
			h.log.WarnContext(r.Context(), "email already used", "error", err)
			response.Error(w, http.StatusConflict, "email already used", err)
			return
		}

		h.log.WarnContext(r.Context(), "failed to update user", "error", err)
		response.Error(w, http.StatusInternalServerError, "failed to update user", err)
		return
	}

	h.log.InfoContext(r.Context(), "success update user", "user_id", id)
	response.Success(w, http.StatusOK, "success update user", user)
}

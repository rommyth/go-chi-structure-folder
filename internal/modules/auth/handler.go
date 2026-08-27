package auth

import (
	"encoding/json"
	"errors"
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
	return &Handler{service: service, log: log, validate: validate}
}

// SignUp godoc
// @Summary 	Sign Up
// @Produce 	json
// @Tags 		Auth
// @Param		request body RegisterRequest true "signup body"
// @Success		200 {object} response.Response
// @Router		/auth/signup [post]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.WarnContext(r.Context(), "invalid request body", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		h.log.WarnContext(r.Context(), "invalid validation", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid validation", err)
		return
	}

	if err := h.service.Register(r.Context(), req); err != nil {
		if errors.Is(err, ErrUserAlreadyExist) {
			h.log.WarnContext(r.Context(), "email already exist", "user_email", req.Email, "error", err)
			response.Error(w, http.StatusConflict, "email already exists", err)
			return
		}

		h.log.WarnContext(r.Context(), "failed to signup", "error", err)
		response.Error(w, http.StatusInternalServerError, "failed to signup", err)
		return
	}

	h.log.InfoContext(r.Context(), "success signup user", "user_email", req.Email)
	response.Success(w, http.StatusCreated, "success signup user", nil)
	return
}

// SignUp godoc
// @Summary 	Login
// @Produce 	json
// @Tags 		Auth
// @Param		request body LoginRequest true "login body"
// @Success		200 {object} response.Response
// @Router		/auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.WarnContext(r.Context(), "invalid request body", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		h.log.WarnContext(r.Context(), "invalid validation", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid validation", err)
		return
	}

	result, err := h.service.Login(r.Context(), req)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			h.log.WarnContext(r.Context(), "invalid email or password", "error", err)
			response.Error(w, http.StatusNotFound, "invalid email or password", err)
			return
		}

		h.log.WarnContext(r.Context(), "failed login", "user_email", req.Email, "error", err)
		response.Error(w, http.StatusUnauthorized, "invalid email or password", err)
		return
	}

	h.log.InfoContext(r.Context(), "success login", "user_email", result.User.Email)
	response.Success(w, http.StatusOK, "success login", result)
}

// Refresh Token godoc
// @Summary 	Refresh Token
// @Produce 	json
// @Tags 		Auth
// @Param		request body RefreshTokenRequest true "refrehtoken body"
// @Success		200 {object} response.Response
// @Router		/auth/refresh-token [post]
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.WarnContext(r.Context(), "invalid request body", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	if err := h.validate.Struct(req); err != nil {
		h.log.WarnContext(r.Context(), "invalid validation", "error", err)
		response.Error(w, http.StatusBadRequest, "invalid validation", err)
		return
	}

	result, err := h.service.RefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		h.log.WarnContext(r.Context(), "invalid refresh token", "error", err)
		response.Error(w, http.StatusUnauthorized, "invalid refresh token", err)
		return
	}

	h.log.WarnContext(r.Context(), "success refresh token", "error", err)
	response.Success(w, http.StatusOK, "success refresh token", result)
}

package user

import (
	"log/slog"
	"net/http"

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

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TODO
}

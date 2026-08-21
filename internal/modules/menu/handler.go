package menu

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

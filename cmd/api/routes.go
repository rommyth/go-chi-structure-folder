package main

import (
	"restaurant-management/internal/modules/health"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Routes struct {
	healthHandler *health.Handler
}

func (r *Routes) LoadRoutes() *chi.Mux {
	route := chi.NewRouter()

	route.Use(middleware.RequestID)
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)

	route.Get("/health", r.healthHandler.CheckHealth)

	return route
}

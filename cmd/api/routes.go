package main

import (
	"net/http"
	authMW "restaurant-management/internal/middleware"
	"restaurant-management/internal/modules/health"
	"restaurant-management/internal/modules/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Routes struct {
	jwt           *jwtauth.JWTAuth
	healthHandler *health.Handler
	userHandler   *user.Handler
}

func (r *Routes) LoadRoutes() *chi.Mux {
	route := chi.NewRouter()

	route.Use(middleware.RequestID)
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)

	// Main Route
	route.Route("/api", func(route chi.Router) {
		route.Group(r.SetupGuestRoute)
		route.Group(r.SetupAuthRoute)
	})

	return route
}

func (r *Routes) SetupGuestRoute(route chi.Router) {
	route.Get("/health", r.healthHandler.CheckHealth)

	route.Get("/guest", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Guest Route"))
	})
}

func (r *Routes) SetupAuthRoute(route chi.Router) {
	route.Use(authMW.Verify(r.jwt))
	route.Use(authMW.Authenticator)

	route.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected Route"))
	})

	user.RegisterRoutes(route, r.userHandler)
}

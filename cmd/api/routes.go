package main

import (
	"fmt"
	"net/http"
	authMW "restaurant-management/internal/middleware"
	"restaurant-management/internal/modules/auth"
	"restaurant-management/internal/modules/food"
	"restaurant-management/internal/modules/health"
	"restaurant-management/internal/modules/menu"
	"restaurant-management/internal/modules/user"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Routes struct {
	jwt           *jwtauth.JWTAuth
	healthHandler *health.Handler
	authHandler   *auth.Handler
	userHandler   *user.Handler
	foodHandler   *food.Handler
	menuHandler   *menu.Handler
}

func (r *Routes) LoadRoutes() *chi.Mux {
	// Initialize Route Engine
	route := chi.NewRouter()

	// Middleware
	route.Use(middleware.RequestID)
	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)

	// Scalar Docs
	route.Route("/docs", func(route chi.Router) {
		route.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./docs/swagger.json")
		})
		route.Get("/", func(w http.ResponseWriter, r *http.Request) {
			htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
				SpecURL:  "http://localhost:3000/docs/swagger.json",
				DarkMode: true,
			})

			if err != nil {
				fmt.Printf("%v", err)
			}

			fmt.Fprintln(w, htmlContent)
		})
	})

	// Main Route
	route.Route("/api", func(route chi.Router) {
		route.Group(r.SetupGuestRoute)
		route.Group(r.SetupProtectedRoute)
	})

	return route
}

func (r *Routes) SetupGuestRoute(route chi.Router) {
	route.Get("/health", r.healthHandler.CheckHealth)

	auth.RegisterRoutes(route, r.authHandler)
}

func (r *Routes) SetupProtectedRoute(route chi.Router) {
	route.Use(authMW.Verify(r.jwt))
	route.Use(authMW.Authenticator)

	user.RegisterRoutes(route, r.userHandler)
	food.RegisterRoutes(route, r.foodHandler)
	menu.RegisterRoutes(route, r.menuHandler)
}

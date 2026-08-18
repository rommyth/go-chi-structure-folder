package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"restaurant-management/internal/modules/health"
	"restaurant-management/internal/modules/user"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type App struct {
	config   *viper.Viper
	db       *pgxpool.Pool
	log      *slog.Logger
	validate *validator.Validate
	jwt      *jwtauth.JWTAuth
}

func NewApp(
	config *viper.Viper,
	log *slog.Logger,
	validate *validator.Validate,
	db *pgxpool.Pool,
	jwt *jwtauth.JWTAuth,
) *App {
	return &App{
		config:   config,
		log:      log,
		validate: validate,
		db:       db,
		jwt:      jwt,
	}
}

// Get All Handler Here
func (a *App) AmountHandler() http.Handler {
	// Initialize Repository
	userRepository := user.NewRepository(a.db)

	// Initilaize Service
	userService := user.NewService(userRepository)

	// Initilaize Handler
	healthHandler := health.NewHandler()
	userHandler := user.NewHandler(userService)

	// Add to root config initialize routes.go
	routesConfig := Routes{
		jwt:           a.jwt,
		healthHandler: healthHandler,
		userHandler:   userHandler,
	}

	return routesConfig.LoadRoutes()
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.config.GetString("server.port")),
		Handler: a.AmountHandler(),
	}

	fmt.Println("Starting Server")

	ch := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			ch <- fmt.Errorf("failed to start server: %w \n", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		fmt.Println("Logging Off")
		shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	}
}

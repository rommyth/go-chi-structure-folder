package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"restaurant-management/internal/modules/auth"
	"restaurant-management/internal/modules/food"
	"restaurant-management/internal/modules/health"
	"restaurant-management/internal/modules/menu"
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
	foodRepository := food.NewRepository(a.db)
	menuRepository := menu.NewRepository(a.db)

	// Initilaize Service
	authService := auth.NewService(userRepository, a.jwt)
	userService := user.NewService(userRepository)
	foodService := food.NewService(foodRepository, menuRepository)
	menuService := menu.NewService(menuRepository)

	// Initilaize Handler
	healthHandler := health.NewHandler()
	authHandler := auth.NewHandler(authService, a.log, a.validate)
	userHandler := user.NewHandler(userService, a.log, a.validate)
	foodHandler := food.NewHandler(foodService, a.log, a.validate)
	menuHandler := menu.NewHandler(menuService, a.log, a.validate)

	// Add to root config initialize routes.go
	routesConfig := Routes{
		jwt:           a.jwt,
		healthHandler: healthHandler,
		authHandler:   authHandler,
		userHandler:   userHandler,
		foodHandler:   foodHandler,
		menuHandler:   menuHandler,
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

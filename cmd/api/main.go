package main

import (
	"context"
	"os"
	"os/signal"
	"restaurant-management/internal/config"
)

// @title		Restaurant Management API
// @version		1.0
// @description	restaurant description
// @host		localhost:3000
// @BasePath	/api
func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger()
	validate := config.NewValidator()
	db, err := config.NewDatabase(viperConfig)
	if err != nil {
		log.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}
	jwt := config.NewJWT(viperConfig)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := NewApp(viperConfig, log, validate, db, jwt)
	if err := app.Start(ctx); err != nil {
		log.Error(
			"application stopped",
			"error", err,
		)

		os.Exit(1)
	}

}

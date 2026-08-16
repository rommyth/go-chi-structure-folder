package main

import (
	"os"
	"restaurant-management/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger()
	validate := config.NewValidator()
	db, err := config.NewDatabase(viperConfig)
	if err != nil {
		log.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}

	app := NewApp(viperConfig, log, validate, db)
	if err := app.Start(); err != nil {
		log.Error(
			"application stopped",
			"error", err,
		)

		os.Exit(1)
	}

}

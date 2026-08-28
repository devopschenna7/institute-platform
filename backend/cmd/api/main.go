package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"institute-platform/backend/internal/bootstrap"
	"institute-platform/backend/internal/config"
	"institute-platform/backend/migrations"
	"institute-platform/backend/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	appLogger := logger.New()

	if err := migrations.Run(cfg.Database); err != nil {
		log.Fatal(err)
	}

	app, err := bootstrap.New(
		ctx,
		cfg,
		appLogger,
	)
	if err != nil {
		appLogger.Error(
			"failed to bootstrap application",
			"error", err,
		)

		os.Exit(1)
	}

	appLogger.Info(
		"starting application",
		"environment", cfg.App.Environment,
		"port", cfg.Server.Port,
	)

	if err := app.Server.Start(); err != nil {
		appLogger.Error(
			"server stopped",
			"error", err,
		)

		os.Exit(1)
	}
}

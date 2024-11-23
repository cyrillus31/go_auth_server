package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/cyrillus31/go_auth_server/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// INFO: инициализировать объект конфига
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// INFO: инициализировать логгер
	log := setupLogger(cfg.Env)

	log.Info(
		"Starting application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPC.Port),
	)
	
	log.Debug("Debug message")
	log.Error("Error message")
	log.Warn("Warning message")

	// TODO: инициализировать приложоения-сервис (app)

	// TODO: запустить gRPC сервер
}

func setupLogger(envVar string) *slog.Logger {
	var log *slog.Logger
	switch envVar {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

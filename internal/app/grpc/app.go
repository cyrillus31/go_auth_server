package app

import (
	"log/slog"

	"google.golang.org/grpc"
	authgrpc "sso/internal/grpc/auth"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

func New(log *slog.Logger, port string) *App {
	gRPCServer := grpc.NewServer()
	return &App{}
}

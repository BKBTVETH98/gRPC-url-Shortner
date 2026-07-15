package app

import (
	"log/slog"
	"time"

	grpcapp "sso/server/internal/app/grpc"
)

type App struct {
	GRPC *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {

	//TODO: инициализировать хранилище
	//TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort, storagePath, tokenTTL)

	return &App{
		GRPC: grpcApp,
	}
}

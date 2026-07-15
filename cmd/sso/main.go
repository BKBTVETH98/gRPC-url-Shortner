package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/server/internal/app"
	"sso/server/internal/config"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO: инициализация объекта конфига

	cfg := config.MustLoad()

	//TODO: инициализация логгера
	log := setupLogger(cfg.Env)

	application := app.New(
		log,
		cfg.GRPC.Port,
		cfg.StoragePath,
		cfg.TokenTTL,
	)

	application.GRPC.MustRun()
	//TODO: инициализация app

	//TODO: запуск app

	//TODO: graceful shutdown

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,пп
		syscall.SIGTERM,
	)

	<-stop // ждём Ctrl+C или SIGTERM

	if err := application.GRPC.Stop(); err != nil {
		log.Error("failed to gracefully shutdown gRPC server", "error", err)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env { //setup sloger based from env
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

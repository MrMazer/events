package main

import (
	"auth/internal/app"
	"auth/internal/config"
	"log/slog"
	"os"
)

func main() {
	cfg := config.Config{}
	config.ConfigInit(&cfg)

	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	app.RunApp(log, cfg)
}

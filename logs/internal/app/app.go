package app

import (
	"logs/internal/config"
	"logs/internal/database"
	repLogs "logs/internal/repository"
	"logs/internal/router"
	serviceLogs "logs/internal/service"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log/slog"
)

func RunApp(log *slog.Logger, cfg config.Config) {
	log.Info("starting app")
	dsn := cfg.DataBasePath
	conn := database.ConnectDB(log, dsn)
	defer conn.Close(context.Background())

	// Используем NewLocationsRepository, а не NewLocationRepository
	LogsRepo := repLogs.NewLogRepository(conn)
	LogsService := serviceLogs.NewLogService(LogsRepo)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.InitRouter(LogsService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

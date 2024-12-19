package app

import (
	"context"
	"locations/internal/config"
	"locations/internal/database"
	repLocations "locations/internal/repository"
	"locations/internal/router"
	serviceLocations "locations/internal/service"
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

	// Используем NewLocationsRepository
	locationsRepo := repLocations.NewLocationsRepository(conn)
	locationsService := serviceLocations.NewLocationService(locationsRepo)

	e := echo.New()

	// Логирование запросов
	e.Use(middleware.Logger())

	// Восстановление от паник
	e.Use(middleware.Recover())

	// Настройка CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"}, // Разрешает все домены (в продакшене укажите конкретные)
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE}, // Разрешённые методы
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization}, // Разрешённые заголовки
		AllowCredentials: true, // Разрешение на отправку куков
	}))

	// Инициализация маршрутов
	router.InitRouter(locationsService, e)

	// Запуск сервера
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

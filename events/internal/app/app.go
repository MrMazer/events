package app

import (
	"context"
	"events/internal/config"
	"events/internal/database"
	repEvents "events/internal/repository/events"
	"events/internal/router"
	serviceEvents "events/internal/service/events"
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

	eventRepo := repEvents.NewEventRepository(conn)
	eventService := serviceEvents.NewEventService(eventRepo)

	e := echo.New()

	// Логирование запросов
	e.Use(middleware.Logger())

	// Восстановление от паник
	e.Use(middleware.Recover())

	// Настроим CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},  // Разрешает все домены (ограничьте в продакшене)
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},  // Разрешаем только нужные методы
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization},  // Разрешаем заголовки Content-Type и Authorization
		AllowCredentials: true,  // Разрешаем использование cookies и авторизации
	}))

	// Инициализация маршрутов
	router.InitRouter(eventService, e)

	// Запуск сервера
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

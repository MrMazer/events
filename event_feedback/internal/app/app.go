package app

import (
	"event_feedback/internal/config"
	"event_feedback/internal/database"
	repEventFeedback "event_feedback/internal/repository"
	"event_feedback/internal/router"
	serviceEventFeedback "event_feedback/internal/service"
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

	eventFeedbackRepo := repEventFeedback.NewEventFeedbackRepository(conn)
	eventFeedbackService := serviceEventFeedback.NewEventFeedbackService(eventFeedbackRepo)

	e := echo.New()

	// Логирование запросов
	e.Use(middleware.Logger())

	// Восстановление от паник
	e.Use(middleware.Recover())

	// Настройка CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"}, // Разрешает все домены (измените на нужные для продакшена)
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE}, // Разрешенные методы
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization}, // Разрешенные заголовки
		AllowCredentials: true, // Разрешение на отправку куков
	}))

	// Инициализация маршрутов
	router.InitRouter(eventFeedbackService, e)

	// Запуск сервера
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
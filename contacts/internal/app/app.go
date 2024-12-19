package app

import (
	"contacts/internal/config"
	"contacts/internal/database"
	repContacts "contacts/internal/repository/contacts"
	"contacts/internal/router"
	serviceContacts "contacts/internal/service/contacts"
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

	contactRepo := repContacts.NewContactRepository(conn)
	contactService := serviceContacts.NewContactService(contactRepo)

	e := echo.New()

	// Логирование запросов
	e.Use(middleware.Logger())

	// Восстановление от паник
	e.Use(middleware.Recover())

	// Настроим CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},  // Разрешает все домены (для разработки можно ограничить конкретными)
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},  // Разрешаем определенные HTTP методы
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization},  // Разрешаем нужные заголовки
		AllowCredentials: true,  // Разрешаем использование куки
	}))

	// Инициализация маршрутов
	router.InitRouter(contactService, e)

	// Запуск сервера
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

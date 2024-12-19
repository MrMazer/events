package router

import (
	"logs/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRouter(logService service.LogService, e *echo.Echo) {

	e.POST("/logs", logService.CreateLog)
	e.GET("/logs/:eventID", logService.GetLogsByEventID)
}
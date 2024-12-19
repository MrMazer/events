package router

import (
	"event_media/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRouter (eventMediaService service.EventMediaService, e *echo.Echo) {
	

	e.POST("/event-media", eventMediaService.AddMedia)
	e.GET("/event-media/:eventID", eventMediaService.GetMediaByEventID)

}

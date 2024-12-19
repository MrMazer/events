package router

import (
	"event_feedback/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRouter(eventFeedbackService service.EventFeedbackService, e *echo.Echo) {
	e.POST("/event-feedback", eventFeedbackService.CreateFeedback)
	e.GET("/event-feedback/:id", eventFeedbackService.GetFeedbackByEventID)

}
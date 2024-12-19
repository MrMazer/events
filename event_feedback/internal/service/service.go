package service

import (
	"event_feedback/internal/models"
	"event_feedback/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

)

type EventFeedbackService interface {
	CreateFeedback(c echo.Context) error
	GetFeedbackByEventID(c echo.Context) error
}

type eventFeedbackService struct {
	repo repository.EventFeedbackRepository
}

func NewEventFeedbackService(repo repository.EventFeedbackRepository) EventFeedbackService {
	return &eventFeedbackService{repo: repo}
}

func (s *eventFeedbackService) CreateFeedback(c echo.Context) error {
	feedback := new(models.EventFeedback)
	if err := c.Bind(feedback); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	id, err := s.repo.CreateFeedback(*feedback)
	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}
	return c.JSON(http.StatusCreated, id)
}

func (s *eventFeedbackService) GetFeedbackByEventID(c echo.Context) error {
    eventID := c.Param("id")
    id, err := strconv.Atoi(eventID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid event ID. Event ID must be a number.",
        })
    }

    feedbacks, err := s.repo.GetFeedbackByEventID(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to fetch feedbacks: " + err.Error(),
        })
    }

    return c.JSON(http.StatusOK, feedbacks)
}


package service

import (
	"event_media/internal/models"
	"event_media/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv" // Импортируем strconv для преобразования типов
)

type EventMediaService interface {
	AddMedia(c echo.Context) error
	GetMediaByEventID(c echo.Context) error
}

type eventMediaService struct {
	repo repository.EventMediaRepository
}

func NewEventMediaService(repo repository.EventMediaRepository) EventMediaService {
	return &eventMediaService{repo: repo}
}

func (s *eventMediaService) AddMedia(c echo.Context) error {
	media := new(models.EventMedia)
	if err := c.Bind(media); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	id, err := s.repo.AddMedia(*media)
	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}
	return c.JSON(http.StatusCreated, id)
}

func (s *eventMediaService) GetMediaByEventID(c echo.Context) error {
	eventID := c.Param("eventID")

	// Преобразуем eventID из строки в int
	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid eventID")
	}

	mediaItems, err := s.repo.GetMediaByEventID(eventIDInt) // Передаем int значение
	if err != nil {
		return c.JSON(http.StatusNotFound, "Media not found")
	}
	return c.JSON(http.StatusOK, mediaItems)
}

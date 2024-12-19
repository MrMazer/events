package service

import (
	"logs/internal/models"
	"logs/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type LogService interface {
	CreateLog(c echo.Context) error
	GetLogsByEventID(c echo.Context) error
}

type logService struct {
	repo repository.LogRepository
}

func NewLogService(repo repository.LogRepository) LogService {
	return &logService{repo: repo}
}

func (s *logService) CreateLog(c echo.Context) error {
	log := new(models.Log)
	if err := c.Bind(log); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	id, err := s.repo.CreateLog(*log)
	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}
	return c.JSON(http.StatusCreated, id)
}

func (s *logService) GetLogsByEventID(c echo.Context) error {
	eventIDStr := c.Param("eventID")

	// Конвертация строки в число
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid event ID")
	}

	// Вызов репозитория
	logs, err := s.repo.GetLogsByEventID(eventID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Logs not found")
	}
	return c.JSON(http.StatusOK, logs)
}

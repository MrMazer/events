package service

import (
	"locations/internal/models"
	"locations/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv" // Импортируем для преобразования строки в int
)

type LocationService interface {
	AddLocation(c echo.Context) error
	GetLocationByID(c echo.Context) error
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) AddLocation(c echo.Context) error {
	location := new(models.Location)
	if err := c.Bind(location); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	id, err := s.repo.AddLocation(*location)
	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}
	return c.JSON(http.StatusCreated, id)
}

func (s *locationService) GetLocationByID(c echo.Context) error {
	locationID := c.Param("locationID")

	// Преобразуем locationID из строки в int
	locationIDInt, err := strconv.Atoi(locationID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid locationID")
	}

	// Передаем преобразованный locationIDInt в репозиторий
	location, err := s.repo.GetLocationByID(locationIDInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Location not found")
	}
	return c.JSON(http.StatusOK, location)
}

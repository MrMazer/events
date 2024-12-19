package router

import (
	"locations/internal/service"
	"github.com/labstack/echo/v4"
)

func InitRouter (locationService service.LocationService, e *echo.Echo) {
	

	
	e.POST("/locations", locationService.AddLocation)
	e.GET("/locations/:locationID", locationService.GetLocationByID)

}

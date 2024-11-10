package controller

import (
	"net/http"

	"github.com/KnoblauchPilze/my-api/pkg/service"
	"github.com/labstack/echo/v4"
)

func GetTemperatureForCity(c echo.Context) error {
	cityName := c.Param("name")

	s := service.NewWeatherService()

	tempData, err := s.GetTemperatureForCity(c.Request().Context(), cityName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetchg temperature")
	}

	return c.JSON(http.StatusOK, tempData)
}

package rest

import (
	"github.com/KnoblauchPilze/weather-app/pkg/controller"
	"github.com/KnoblauchPilze/weather-app/pkg/server"
)

func AddWeatherRoutes(s server.Server) {
	s.GET("/weather/city/:name", controller.GetTemperatureForCity)
}

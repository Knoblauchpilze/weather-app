package rest

import (
	"github.com/KnoblauchPilze/my-api/pkg/controller"
	"github.com/KnoblauchPilze/my-api/pkg/server"
)

func AddWeatherRoutes(s server.Server) {
	s.GET("/weather/city/:name", controller.GetTemperatureForCity)
}

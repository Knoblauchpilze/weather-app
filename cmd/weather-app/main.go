package main

import (
	"github.com/KnoblauchPilze/weather-app/pkg/rest"
	"github.com/KnoblauchPilze/weather-app/pkg/server"
)

func main() {
	s := server.New()

	rest.AddWeatherRoutes(s)

	port := 1234
	s.Start(port)
}

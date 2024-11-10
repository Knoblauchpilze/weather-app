package main

import (
	"github.com/KnoblauchPilze/my-api/pkg/rest"
	"github.com/KnoblauchPilze/my-api/pkg/server"
)

func main() {
	s := server.New()

	rest.AddWeatherRoutes(s)

	port := 1234
	s.Start(port)
}

package rest

import (
	"testing"

	"github.com/KnoblauchPilze/my-api/pkg/server"
	"github.com/stretchr/testify/mock"
)

func Test_AddWeatherRoutes(t *testing.T) {
	m := server.NewMockServer(t)

	m.EXPECT().GET("/weather/city/:name", mock.AnythingOfType("echo.HandlerFunc")).Times(1)

	AddWeatherRoutes(m)
}

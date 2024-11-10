package server

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
)

type Server interface {
	Start(port int) error

	DELETE(path string, h echo.HandlerFunc)
	GET(path string, h echo.HandlerFunc)
	PATCH(path string, h echo.HandlerFunc)
	POST(path string, h echo.HandlerFunc)
}

type serverImpl struct {
	server *echo.Echo
}

func New() Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	log := zerolog.New(os.Stdout)
	e.Logger = lecho.From(log, lecho.WithTimestamp())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &serverImpl{
		server: e,
	}
}

func (s *serverImpl) Start(port int) error {
	address := fmt.Sprintf(":%d", port)

	s.server.Logger.Infof("Starting server on port %d", port)
	return s.server.Start(address)
}

func (s *serverImpl) DELETE(path string, h echo.HandlerFunc) {
	s.server.DELETE(path, h)
}

func (s *serverImpl) GET(path string, h echo.HandlerFunc) {
	s.server.GET(path, h)
}

func (s *serverImpl) PATCH(path string, h echo.HandlerFunc) {
	s.server.PATCH(path, h)
}

func (s *serverImpl) POST(path string, h echo.HandlerFunc) {
	s.server.POST(path, h)
}

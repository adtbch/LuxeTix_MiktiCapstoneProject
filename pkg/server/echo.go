package server

import (
	"luxe/config"
	"luxe/pkg/route"

	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
}

func NewServer(cfg *config.Config, publicRoutes []route.Route) *Server {
	e := echo.New()
	v1 := e.Group("/api/v1")
	if len(publicRoutes) > 0 {
		for _, route := range publicRoutes {
			v1.Add(route.Method, route.Path, route.Handler)
		}
	}
	return &Server{e}
}
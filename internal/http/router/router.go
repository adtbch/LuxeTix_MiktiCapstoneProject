package router

import (
	"luxe/internal/http/handler"
	"luxe/pkg/route"
	"net/http"
)

var (
	adminOnly = []string{"Administrator"}
	allRoles = []string{"Administrator", "User"}
)

func PublicRoutes(userHandler handler.UserHandler) []route.Route {
	return []route.Route{
		{
			Method: http.MethodPost,
			Path: "/register",
			Handler: userHandler.Register,
		},
		{
			Method: http.MethodPost,
			Path: "/login",
			Handler: userHandler.Login,
		},
	}
}
package router

import (
	"net/http"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/handler"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/route"
)

func PublicRoutes(event handler.EventHandler) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/events",
			Handler: event.GetEvents,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/:id",
			Handler: event.GetEvent,
		},
		{
			Method:  http.MethodPost,
			Path:    "/events",
			Handler: event.CreateEvent,
		},
		{
			Method:  http.MethodPut,
			Path:    "/events/:id",
			Handler: event.UpdateEvent,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/events/:id",
			Handler: event.DeleteEvent,
		},
	}
}

func PrivateRoutes(event handler.EventHandler) []route.Route {
	return []route.Route{}
}

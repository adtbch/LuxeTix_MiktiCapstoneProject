package router

import (
	"net/http"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/handler"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/route"
)

func PublicRoutes(event handler.EventHandler, user handler.UserHandler) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/events",
			Handler: event.GetEvents,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: user.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: user.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/request-reset-password",
			Handler: user.ResetPasswordRequest,
		},
		{
			Method:  http.MethodPost,
			Path:    "/reset-password/:token",
			Handler: user.ResetPassword,
		},
		{
			Method:  http.MethodGet,
			Path:    "/verify-email/:token",
			Handler: user.VerifyEmail,
		},	
		{
			Method:  http.MethodGet,
			Path:    "/events/:id",
			Handler: event.GetEvent,
		},
	}
}

func PrivateRoutes(event handler.EventHandler, user handler.UserHandler, tran handler.TransactionHandler, notif handler.NotificationHandler) []route.Route {
	return []route.Route{

		{
			Method:  http.MethodPost,
			Path:    "/events",
			Handler: event.CreateEvent,
		},
		{
			Method:  http.MethodPut,
			Path:    "/events/:id",
			Handler: event.UpdateEventByUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/events/:id",
			Handler: event.DeleteEvent,
		},
		{
			Method:  http.MethodPost,
			Path:    "/create-order",
			Handler: tran.CreateOrder,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: user.GetProfil,
		},
		{ 
			Method:  http.MethodGet,
			Path:    "/notification",
			Handler: notif.UserGetNotification,
		},
	}
}

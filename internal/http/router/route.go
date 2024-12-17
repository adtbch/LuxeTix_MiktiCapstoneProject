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
			Path:    "/events/sort/expensive",
			Handler: event.SortFromExpensivestToCheapest,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/sort/cheapest",
			Handler: event.SortFromCheapestToExpensivest,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/sort/newest",
			Handler: event.SortNewestToOldest,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/filterByCategory/:category",
			Handler: event.FilteringByCategory,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/filterByLocation/:location",
			Handler: event.FilteringByLocation,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/filterByDate/:date",
			Handler: event.FilteringByDate,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/filterByPrice/:price",
			Handler: event.FilteringByPrice,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/filterByMinPriceAndMaxPrice/:min_price/:max_price",
			Handler: event.FilteringMInMaxPrice, // corrected the typo
		},
		{
			Method:  http.MethodGet,
			Path:    "/events/:id",
			Handler: event.GetEvent,
		},
	}
}

func PrivateRoutes(event handler.EventHandler, user handler.UserHandler) []route.Route {
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
	}
}

package router

import (
	"net/http"

	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/handler"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/route"
)

var (
	adminOnly = []string{"Administrator"}
	userOnly  = []string{"User"}
	AllRoles  = []string{"Admninistrator", "User"}
)

func PublicRoutes(event handler.EventHandler, user handler.UserHandler, tran handler.TransactionHandler) []route.Route {
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
		{
			Method:  http.MethodPost,
			Path:    "/webhook",
			Handler: tran.WebHookTransaction,
		},
	}
}

func PrivateRoutes(event handler.EventHandler, user handler.UserHandler, tran handler.TransactionHandler, notif handler.NotificationHandler, submissions handler.RequestEventHandler) []route.Route {
	return []route.Route{
		//event
		//role user
		{
			Method:  http.MethodPut,
			Path:    "/events/:id",
			Handler: event.UpdateEventByUser,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "/events-owner",
			Handler: event.GetAllEventByOwner,
			Roles:   userOnly,
		},
		//role admin
		{
			Method:  http.MethodDelete,
			Path:    "/events/:id",
			Handler: event.DeleteEvent,
			Roles:   adminOnly,
		},

		//Transactions
		{
			Method:  http.MethodPost,
			Path:    "/order",
			Handler: tran.CreateOrder,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "/order",
			Handler: tran.GetUserTransactions,
			Roles:   userOnly,
		},

		//user
		//role user
		{
			Method:  http.MethodGet,
			Path:    "/users/profile",
			Handler: user.GetProfil,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "/users/profile",
			Handler: user.UpdateProfil,
			Roles:   userOnly,
		},

		//role admin
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: user.GetUsers,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/:id",
			Handler: user.GetUser,
			Roles:   adminOnly,
		},

		//notification
		//role user
		{
			Method:  http.MethodGet,
			Path:    "/notification",
			Handler: notif.UserGetNotification,
			Roles:   userOnly,
		},
		//role admin
		{
			Method:  http.MethodGet,
			Path:    "/notification",
			Handler: notif.GetAllNotification,
			Roles:   adminOnly,
		},

		//SUBMISSION
		//role user
		{
			Method:  http.MethodPost,
			Path:    "/submissions",
			Handler: submissions.CreateEvent,
			Roles:   userOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "/submissions",
			Handler: submissions.UpdateSubmissionByUser,
			Roles:   userOnly,
		},
		//role admin
		{
			Method:  http.MethodDelete,
			Path:    "/submissions/:id",
			Handler: submissions.CancelSubmission,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "/submissions/:id/approve",
			Handler: submissions.ApproveSubmission,
			Roles:   adminOnly,
		},
		{
			Method:  http.MethodPut,
			Path:    "/submissions/:id/reject",
			Handler: submissions.RejectSubmission,
			Roles:   adminOnly,
		},
	}
}

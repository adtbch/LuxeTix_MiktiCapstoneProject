package builder

import (
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/config"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/handler"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/router"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
	service "github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/services"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/route"
	"github.com/midtrans/midtrans-go/snap"

	"gorm.io/gorm"
)

func BuilderPublicRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []route.Route {
	//repository
	eventRepository := repository.NewEventRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	userRepository := repository.NewUserRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	//end

	//service
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	notificationService := service.NewNotificationService(notificationRepository)
	userService := service.NewUserService(tokenService, cfg, userRepository)
	eventService := service.NewEventService(eventRepository, transactionRepository)
	paymentService := service.NewPaymentService(midtransClient, cfg, notificationService)
	tranService := service.NewTransactionService(cfg, transactionRepository, eventRepository, userRepository, paymentService)
	//end
	
	//handler
	eventHandler := handler.NewEventHandler(eventService, tokenService)
	tranHandler := handler.NewTransactionHandler(tranService, tokenService, paymentService, userService, eventService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	//end

	return router.PublicRoutes(eventHandler, userHandler, tranHandler)
}

func BuilderPrivateRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []route.Route {
	//repository
	eventRepository := repository.NewEventRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	userRepository := repository.NewUserRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)
	requestEventRepository := repository.NewRequestEventRepository(db)
	//end

	//service
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	userService := service.NewUserService(tokenService, cfg, userRepository)
	eventService := service.NewEventService(eventRepository, transactionRepository)
	notifService := service.NewNotificationService(notificationRepository)
	submissionService := service.NewSubmissionService(cfg, requestEventRepository, transactionRepository, userRepository, notifService)
	paymentService := service.NewPaymentService(midtransClient,cfg, notifService)
	tranService := service.NewTransactionService(cfg, transactionRepository, eventRepository, userRepository, paymentService)

	//end

	//handler
	eventHandler := handler.NewEventHandler(eventService, tokenService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	tranHandler := handler.NewTransactionHandler(tranService, tokenService, paymentService, userService, eventService)
	notifHandler := handler.NewNotificationHandler(notifService, tokenService)
	submissionHandler := handler.NewRequestEventHandler(submissionService, tokenService)
	//end

	return router.PrivateRoutes(eventHandler, userHandler, tranHandler, notifHandler, submissionHandler)
}

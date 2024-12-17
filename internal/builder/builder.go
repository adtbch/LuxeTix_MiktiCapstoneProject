package builder

import (
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/config"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/router"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/repository"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/services"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/internal/http/handler"
	"github.com/adtbch/LuxeTix_MiktiCapstoneProject/pkg/route"
	
	"gorm.io/gorm"
)

func BuilderPublicRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	//repository
	eventRepository := repository.NewEventRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	userRepository := repository.NewUserRepository(db)
	//end

	//service
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	userService := service.NewUserService(tokenService, cfg, userRepository)
	eventService := service.NewEventService(eventRepository, transactionRepository)
	//end

	//handler
	eventHandler := handler.NewEventHandler(eventService, tokenService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	//end

	return router.PublicRoutes(eventHandler, userHandler)
}

func BuilderPrivateRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	//repository
	eventRepository := repository.NewEventRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	userRepository := repository.NewUserRepository(db)
	//end

	//service
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	userService := service.NewUserService(tokenService, cfg, userRepository)
	eventService := service.NewEventService(eventRepository, transactionRepository)
	// _ = service.NewEventService()
	//end

	//handler
	eventHandler := handler.NewEventHandler(eventService, tokenService)
	userHandler := handler.NewUserHandler(tokenService, userService)
	//end

	return router.PrivateRoutes(eventHandler, userHandler)
}

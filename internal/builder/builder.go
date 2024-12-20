package builder

import (
	"luxe/config"
	"luxe/internal/http/handler"
	"luxe/internal/http/router"
	"luxe/internal/repository"
	"luxe/internal/service"
	"luxe/pkg/route"

	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	// repository
	userRepository := repository.NewUserRepository(db)
	
	// service
	tokenService := service.NewTokenService(cfg.JWTConfig.SecretKey)
	userService := service.NewUserService(cfg, userRepository)

	// handler
	userHandler := handler.NewUserHandler(tokenService, userService)

	return router.PublicRoutes(userHandler)
}
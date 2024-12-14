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
	//end

	//service
	eventService := service.NewEventService(eventRepository)
	//end

	//handler
	eventHandler := handler.NewEventHandler(eventService)
	//end

	return router.PublicRoutes(eventHandler)
}

func BuilderPrivateRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	//repository
	_ = repository.NewEventRepository(db)
	//end

	//service
	// _ = service.NewEventService()
	//end

	//handler
	// _ = handler.NewEventHandler()
	//end

	return nil
}

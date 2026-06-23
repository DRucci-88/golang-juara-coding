package routes

import (
	"materi-middleware-gorm/handler"
	"materi-middleware-gorm/middleware"
	"materi-middleware-gorm/repository"
	"materi-middleware-gorm/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Global Middlewares
	r.Use(middleware.ApiKeyAuth())
	r.Use(middleware.RateLimiter(5))

	// Initialize layers
	concertRepo := repository.NewConcertRepository(db)
	concertService := service.NewConcertService(concertRepo)
	concertHandler := handler.NewConcertHandler(concertService)

	ticketCategoryRepo := repository.NewTicketCategoryRepository(db)
	ticketCategoryService := service.NewTicketCategoryService(ticketCategoryRepo, concertRepo)
	ticketCategoryHandler := handler.NewTicketCategoryHandler(ticketCategoryService)

	// Group routes
	api := r.Group("/api/v1")
	{
		// Concerts routes
		api.POST("/concerts", concertHandler.CreateConcert)
		api.GET("/concerts", concertHandler.GetConcerts)
		api.GET("/concerts/:id", concertHandler.GetConcertByID)
		api.PUT("/concerts/:id", concertHandler.UpdateConcert)
		api.DELETE("/concerts/:id", concertHandler.DeleteConcert)

		// Ticket Categories routes
		api.POST("/ticket-categories", ticketCategoryHandler.CreateTicketCategory)
		api.GET("/ticket-categories", ticketCategoryHandler.GetTicketCategories)
		api.GET("/ticket-categories/:id", ticketCategoryHandler.GetTicketCategoryByID)
		api.PUT("/ticket-categories/:id", ticketCategoryHandler.UpdateTicketCategory)
		api.DELETE("/ticket-categories/:id", ticketCategoryHandler.DeleteTicketCategory)
	}

	return r
}

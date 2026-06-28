package app

import (
	"praktikum/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	orderHandler handler.OrderHandler,
	analyticsHandler handler.AnalyticsHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Berjalan Perfecto"})
	})

	api := r.Group("/api/v1")
	api.POST("/orders", orderHandler.Create)
	api.GET("/orders/:id", orderHandler.FindByID)
	api.PUT("/orders/:id/cancel", orderHandler.Cancel)
	api.GET("/analytics/popular-products", analyticsHandler.PopularProduct)

	return r
}

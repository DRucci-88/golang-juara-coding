package app

import (
	"praktikum/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	m *GroupMiddleware,
	authHandler handler.AuthHandler,
	orderHandler handler.OrderHandler,
	analyticsHandler handler.AnalyticsHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Berjalan Perfecto"})
	})

	api := r.Group("/api/v1")
	api.POST("/orders", m.JWT, orderHandler.Create)
	api.GET("/orders/:id", orderHandler.FindByID)
	api.PUT("/orders/:id/cancel", orderHandler.Cancel)
	api.GET("/analytics/popular-products", analyticsHandler.PopularProduct)

	authApi := api.Group("/auth")
	authApi.POST("/register", authHandler.Register)
	authApi.POST("/login", authHandler.Login)
	authApi.GET("/me", m.JWT, authHandler.Me)

	return r
}

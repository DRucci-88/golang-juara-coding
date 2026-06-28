package app

import (
	"praktikum/handler"
	"praktikum/middleware"

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
	api.POST("/orders", m.JWT, middleware.RequiredRoleMiddleware("ADMIN"), orderHandler.Create)
	api.GET("/orders/:id", orderHandler.FindByID)
	api.PUT("/orders/:id/cancel", orderHandler.Cancel)
	api.GET("/analytics/popular-products", analyticsHandler.PopularProduct)

	authApi := api.Group("/auth")
	authApi.POST("/register", authHandler.Register)
	authApi.POST("/login", authHandler.Login)
	authApi.GET("/me-admin", m.JWT, middleware.RequiredRoleMiddleware("ADMIN"), authHandler.Me)
	authApi.GET("/me-user", m.JWT, middleware.RequiredRoleMiddleware("USER"), authHandler.Me)

	return r
}

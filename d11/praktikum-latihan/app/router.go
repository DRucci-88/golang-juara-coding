package app

import (
	"praktikum/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	orderHanler handler.OrderHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Berjalan Perfecto"})
	})

	api := r.Group("/api/v1")
	api.POST("/orders", orderHanler.Create)

	return r
}

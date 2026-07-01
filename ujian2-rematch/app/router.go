package app

import (
	"ujian2_rematch/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	departmentHander *handler.DepartmentHandler,
	positionHandler *handler.PositionHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Berjalan Perfecto"})
	})

	departmentApi := r.Group("/departments")
	departmentApi.POST("", departmentHander.Create)
	departmentApi.GET("/:id", departmentHander.FindByID)
	departmentApi.GET("", departmentHander.FindAll)
	departmentApi.PUT("/:id", departmentHander.Update)
	departmentApi.DELETE("/:id", departmentHander.Delete)

	positionApi := r.Group("/positions")
	positionApi.POST("", positionHandler.Create)
	positionApi.GET("/:id", positionHandler.FindByID)
	positionApi.GET("", positionHandler.FindAll)
	positionApi.PUT("/:id", positionHandler.Update)
	positionApi.DELETE("/:id", positionHandler.Delete)

	return r
}

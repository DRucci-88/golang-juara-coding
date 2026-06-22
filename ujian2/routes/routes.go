package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mini-hris/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "mini-hris service is running"})
	})

	api := router.Group("/api")
	{
		api.POST("/departments", controllers.CreateDepartment)
		api.GET("/departments", controllers.ListDepartments)
		api.PUT("/departments/:id", controllers.UpdateDepartment)
		api.DELETE("/departments/:id", controllers.DeleteDepartment)

		api.POST("/positions", controllers.CreatePosition)
		api.GET("/positions", controllers.ListPositions)
		api.PUT("/positions/:id", controllers.UpdatePosition)
		api.DELETE("/positions/:id", controllers.DeletePosition)

		api.POST("/employees", controllers.CreateEmployee)
		api.GET("/employees", controllers.ListEmployees)
		api.PUT("/employees/:id", controllers.UpdateEmployee)

		api.POST("/attendances", controllers.CreateAttendance)

		api.POST("/leaves", controllers.CreateLeave)
		api.PATCH("/leaves/:id/approve", controllers.ApproveLeave)

		api.POST("/salaries/calculate", controllers.CalculateSalaries)
		api.GET("/salaries/period/:period", controllers.ListSalariesByPeriod)
	}
}

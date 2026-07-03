package app

import (
	"log"
	"ujian2_rematch/dto"
	"ujian2_rematch/handler"
	"ujian2_rematch/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	m *middleware.MiddlewareManager,
	authHandler *handler.AuthHandler,
	departmentHander *handler.DepartmentHandler,
	positionHandler *handler.PositionHandler,
	employeeHandler *handler.EmployeeHandler,
	attendanceHandler *handler.AttendanceHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Berjalan Perfecto"})
	})

	authApi := r.Group("/auth")
	authApi.GET("/jwt", m.JWT, func(ctx *gin.Context) {
		log.Println("auth / jwt")
		authContext, exist := ctx.Get("auth")
		if !exist {
			ctx.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		auth := authContext.(*dto.AuthContext)
		ctx.JSON(200, gin.H{"data": auth})
	})
	authApi.POST("/login", authHandler.Login)
	authApi.POST("/logout", m.JWT, authHandler.Logout)

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

	employeeApi := r.Group("/employees")
	employeeApi.POST("", employeeHandler.Create)
	employeeApi.GET("/:id", employeeHandler.FindByID)
	employeeApi.GET("", employeeHandler.FindAll)
	employeeApi.PUT("/:id", employeeHandler.Update)

	attendanceApi := r.Group("/attendances", m.JWT)
	attendanceApi.POST("/check-in", attendanceHandler.CheckIn)
	attendanceApi.POST("/check-out", attendanceHandler.CheckOut)

	return r
}

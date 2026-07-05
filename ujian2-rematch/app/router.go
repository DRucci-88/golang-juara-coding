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
	auth *handler.AuthHandler,
	department *handler.DepartmentHandler,
	position *handler.PositionHandler,
	employee *handler.EmployeeHandler,
	attendance *handler.AttendanceHandler,
	leave *handler.LeaveHandler,
	salary *handler.SalaryHandler,
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
	authApi.POST("/login", auth.Login)
	authApi.POST("/logout", m.JWT, auth.Logout)

	departmentApi := r.Group("/departments")
	departmentApi.POST("", department.Create)
	departmentApi.GET("/:id", department.FindByID)
	departmentApi.GET("", department.FindAll)
	departmentApi.PUT("/:id", department.Update)
	departmentApi.DELETE("/:id", department.Delete)

	positionApi := r.Group("/positions")
	positionApi.POST("", position.Create)
	positionApi.GET("/:id", position.FindByID)
	positionApi.GET("", position.FindAll)
	positionApi.PUT("/:id", position.Update)
	positionApi.DELETE("/:id", position.Delete)

	employeeApi := r.Group("/employees")
	employeeApi.POST("", employee.Create)
	employeeApi.GET("/:id", employee.FindByID)
	employeeApi.GET("", employee.FindAll)
	employeeApi.PUT("/:id", employee.Update)

	attendanceApi := r.Group("/attendances", m.JWT)
	attendanceApi.POST("/check-in", attendance.CheckIn)
	attendanceApi.POST("/check-out", attendance.CheckOut)

	leaveApi := r.Group("/leaves", m.JWT)
	leaveApi.POST("", leave.Create)
	leaveApi.POST("/:id/approve", leave.Approval)

	salaryApi := r.Group("/salaries", m.JWT)
	salaryApi.POST("/calculate", salary.Calculate)
	return r
}

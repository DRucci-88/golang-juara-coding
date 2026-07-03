package handler

import (
	"net/http"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/service"

	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	attendanceService *service.AttendanceService
}

func NewAttendanceHandler(
	attendanceService *service.AttendanceService,
) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
	}
}

func (h *AttendanceHandler) CheckIn(c *gin.Context) {
	authContextValue, exist := c.Get("auth")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": model.ErrAuthUnauthorized,
		})
		return
	}

	authContext := authContextValue.(*dto.AuthContext)
	var employeeID uint = 0
	if authContext.EmployeeID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": model.ErrAuthUnauthorized,
		})
		return
	}
	employeeID = *authContext.EmployeeID

	attendance, err := h.attendanceService.CheckIn(c.Request.Context(), employeeID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewAttendanceResponse(attendance)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *AttendanceHandler) CheckOut(c *gin.Context) {
	authContextValue, exist := c.Get("auth")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": model.ErrAuthUnauthorized,
		})
		return
	}

	authContext := authContextValue.(*dto.AuthContext)
	var employeeID uint = 0
	if authContext.EmployeeID == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": model.ErrAuthUnauthorized,
		})
		return
	}
	employeeID = *authContext.EmployeeID

	attendance, err := h.attendanceService.CheckOut(c.Request.Context(), employeeID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewAttendanceResponse(attendance)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

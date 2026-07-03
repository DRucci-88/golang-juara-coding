package handler

import (
	"net/http"
	"strconv"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/service"

	"github.com/gin-gonic/gin"
)

type LeaveHandler struct {
	leaveService *service.LeaveService
}

func NewLeaveHandler(
	leaveService *service.LeaveService,
) *LeaveHandler {
	return &LeaveHandler{
		leaveService: leaveService,
	}
}

func (h *LeaveHandler) Create(c *gin.Context) {
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

	var req dto.LeaveCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	leave, err := h.leaveService.Create(c.Request.Context(), employeeID, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewLeaveResponse(leave)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

func (h *LeaveHandler) Approval(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	var req dto.LeaveApprovalRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	leave, err := h.leaveService.Approval(c.Request.Context(), id, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.NewLeaveResponse(leave)
	c.JSON(http.StatusCreated, gin.H{"data": res})
}

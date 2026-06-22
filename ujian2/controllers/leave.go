package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mini-hris/config"
	"mini-hris/models"
)

type leaveRequest struct {
	EmployeeID uint   `json:"employee_id" binding:"required,gt=0"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
	Reason     string `json:"reason" binding:"required"`
}

type leaveApprovalRequest struct {
	Status string `json:"status" binding:"required,oneof=APPROVED REJECTED"`
}

func CreateLeave(ctx *gin.Context) {
	var req leaveRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ensureEmployeeExists(req.EmployeeID); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	startDate, err := parseDate(req.StartDate)
	if err != nil {
		respondError(ctx, http.StatusBadRequest, "start_date must use YYYY-MM-DD format")
		return
	}

	endDate, err := parseDate(req.EndDate)
	if err != nil {
		respondError(ctx, http.StatusBadRequest, "end_date must use YYYY-MM-DD format")
		return
	}

	req.Reason = normalizeText(req.Reason)
	if req.Reason == "" {
		respondError(ctx, http.StatusBadRequest, "reason is required")
		return
	}

	if endDate.Before(startDate) {
		respondError(ctx, http.StatusBadRequest, "end_date must be later than or equal to start_date")
		return
	}

	leave := models.Leave{
		EmployeeID: req.EmployeeID,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
		Reason:     req.Reason,
		Status:     "PENDING",
	}

	if err := config.DB.Create(&leave).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Preload("Employee").First(&leave, leave.ID).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "leave created",
		"data":    leave,
	})
}

func ApproveLeave(ctx *gin.Context) {
	id, ok := parseUintParam(ctx, "id")
	if !ok {
		return
	}

	var leave models.Leave
	if err := config.DB.First(&leave, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(ctx, http.StatusNotFound, "leave not found")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var req leaveApprovalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	leave.Status = req.Status
	if err := config.DB.Save(&leave).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Preload("Employee").First(&leave, leave.ID).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "leave status updated",
		"data":    leave,
	})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mini-hris/config"
	"mini-hris/models"
)

type attendanceRequest struct {
	EmployeeID uint   `json:"employee_id" binding:"required,gt=0"`
	Date       string `json:"date" binding:"required"`
	CheckIn    string `json:"check_in" binding:"required"`
	CheckOut   string `json:"check_out" binding:"required"`
	Status     string `json:"status" binding:"required,oneof=PRESENT LATE ABSENT"`
}

func CreateAttendance(ctx *gin.Context) {
	var req attendanceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ensureEmployeeExists(req.EmployeeID); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := parseDate(req.Date); err != nil {
		respondError(ctx, http.StatusBadRequest, "date must use YYYY-MM-DD format")
		return
	}

	checkIn, err := parseClock(req.CheckIn)
	if err != nil {
		respondError(ctx, http.StatusBadRequest, "check_in must use HH:MM format")
		return
	}

	checkOut, err := parseClock(req.CheckOut)
	if err != nil {
		respondError(ctx, http.StatusBadRequest, "check_out must use HH:MM format")
		return
	}

	if checkOut.Before(checkIn) {
		respondError(ctx, http.StatusBadRequest, "check_out must be later than or equal to check_in")
		return
	}

	attendance := models.Attendance{
		EmployeeID: req.EmployeeID,
		Date:       req.Date,
		CheckIn:    req.CheckIn,
		CheckOut:   req.CheckOut,
		Status:     req.Status,
	}

	if err := config.DB.Create(&attendance).Error; err != nil {
		if isUniqueConstraintError(err) {
			respondError(ctx, http.StatusConflict, "attendance for this employee and date already exists")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Preload("Employee").First(&attendance, attendance.ID).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "attendance created",
		"data":    attendance,
	})
}

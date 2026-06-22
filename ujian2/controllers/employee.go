package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mini-hris/config"
	"mini-hris/models"
)

type employeeRequest struct {
	NIK          string `json:"nik" binding:"required"`
	FullName     string `json:"full_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID uint   `json:"department_id" binding:"required,gt=0"`
	PositionID   uint   `json:"position_id" binding:"required,gt=0"`
	Status       string `json:"status" binding:"required,oneof=ACTIVE SUSPENDED TERMINATED"`
	LeaveBalance *int   `json:"leave_balance" binding:"omitempty,min=0"`
}

func CreateEmployee(ctx *gin.Context) {
	var req employeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.NIK = normalizeText(req.NIK)
	req.FullName = normalizeText(req.FullName)
	req.Email = normalizeText(req.Email)
	if req.NIK == "" || req.FullName == "" || req.Email == "" {
		respondError(ctx, http.StatusBadRequest, "nik, full_name, and email are required")
		return
	}

	if err := ensureDepartmentExists(req.DepartmentID); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ensurePositionExists(req.PositionID); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	leaveBalance := 12
	if req.LeaveBalance != nil {
		leaveBalance = *req.LeaveBalance
	}

	employee := models.Employee{
		NIK:          req.NIK,
		FullName:     req.FullName,
		Email:        req.Email,
		DepartmentID: req.DepartmentID,
		PositionID:   req.PositionID,
		Status:       req.Status,
		LeaveBalance: leaveBalance,
	}

	if err := config.DB.Create(&employee).Error; err != nil {
		if isUniqueConstraintError(err) {
			respondError(ctx, http.StatusConflict, "employee nik or email already exists")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Preload("Department").Preload("Position").First(&employee, employee.ID).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "employee created",
		"data":    employee,
	})
}

func ListEmployees(ctx *gin.Context) {
	query := config.DB.Preload("Department").Preload("Position").Model(&models.Employee{})

	if search := strings.TrimSpace(ctx.Query("search")); search != "" {
		query = query.Where("LOWER(full_name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	if departmentID := strings.TrimSpace(ctx.Query("department_id")); departmentID != "" {
		value, err := strconv.Atoi(departmentID)
		if err != nil || value <= 0 {
			respondError(ctx, http.StatusBadRequest, "invalid department_id")
			return
		}
		query = query.Where("department_id = ?", value)
	}

	if positionID := strings.TrimSpace(ctx.Query("position_id")); positionID != "" {
		value, err := strconv.Atoi(positionID)
		if err != nil || value <= 0 {
			respondError(ctx, http.StatusBadRequest, "invalid position_id")
			return
		}
		query = query.Where("position_id = ?", value)
	}

	if status := strings.TrimSpace(ctx.Query("status")); status != "" {
		switch status {
		case "ACTIVE", "SUSPENDED", "TERMINATED":
			query = query.Where("status = ?", status)
		default:
			respondError(ctx, http.StatusBadRequest, "invalid status filter")
			return
		}
	}

	var employees []models.Employee
	if err := query.Order("id ASC").Find(&employees).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": employees})
}

func UpdateEmployee(ctx *gin.Context) {
	id, ok := parseUintParam(ctx, "id")
	if !ok {
		return
	}

	var employee models.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(ctx, http.StatusNotFound, "employee not found")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var req employeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.NIK = normalizeText(req.NIK)
	req.FullName = normalizeText(req.FullName)
	req.Email = normalizeText(req.Email)
	if req.NIK == "" || req.FullName == "" || req.Email == "" {
		respondError(ctx, http.StatusBadRequest, "nik, full_name, and email are required")
		return
	}

	if err := ensureDepartmentExists(req.DepartmentID); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := ensurePositionExists(req.PositionID); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	employee.NIK = req.NIK
	employee.FullName = req.FullName
	employee.Email = req.Email
	employee.DepartmentID = req.DepartmentID
	employee.PositionID = req.PositionID
	employee.Status = req.Status
	if req.LeaveBalance != nil {
		employee.LeaveBalance = *req.LeaveBalance
	}

	if err := config.DB.Save(&employee).Error; err != nil {
		if isUniqueConstraintError(err) {
			respondError(ctx, http.StatusConflict, "employee nik or email already exists")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Preload("Department").Preload("Position").First(&employee, employee.ID).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "employee updated",
		"data":    employee,
	})
}

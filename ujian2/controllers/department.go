package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mini-hris/config"
	"mini-hris/models"
)

type departmentRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

func CreateDepartment(ctx *gin.Context) {
	var req departmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.Name = normalizeText(req.Name)
	req.Code = normalizeText(req.Code)
	if req.Name == "" || req.Code == "" {
		respondError(ctx, http.StatusBadRequest, "name and code are required")
		return
	}

	department := models.Department{
		Name: req.Name,
		Code: req.Code,
	}

	if err := config.DB.Create(&department).Error; err != nil {
		if isUniqueConstraintError(err) {
			respondError(ctx, http.StatusConflict, "department code already exists")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "department created",
		"data":    department,
	})
}

func ListDepartments(ctx *gin.Context) {
	var departments []models.Department
	if err := config.DB.Order("id ASC").Find(&departments).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": departments})
}

func UpdateDepartment(ctx *gin.Context) {
	id, ok := parseUintParam(ctx, "id")
	if !ok {
		return
	}

	var department models.Department
	if err := config.DB.First(&department, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(ctx, http.StatusNotFound, "department not found")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var req departmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.Name = normalizeText(req.Name)
	req.Code = normalizeText(req.Code)
	if req.Name == "" || req.Code == "" {
		respondError(ctx, http.StatusBadRequest, "name and code are required")
		return
	}

	department.Name = req.Name
	department.Code = req.Code

	if err := config.DB.Save(&department).Error; err != nil {
		if isUniqueConstraintError(err) {
			respondError(ctx, http.StatusConflict, "department code already exists")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "department updated",
		"data":    department,
	})
}

func DeleteDepartment(ctx *gin.Context) {
	id, ok := parseUintParam(ctx, "id")
	if !ok {
		return
	}

	var department models.Department
	if err := config.DB.First(&department, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(ctx, http.StatusNotFound, "department not found")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Delete(&department).Error; err != nil {
		if isForeignKeyConstraintError(err) {
			respondError(ctx, http.StatusConflict, "department is still used by employees")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "department deleted"})
}

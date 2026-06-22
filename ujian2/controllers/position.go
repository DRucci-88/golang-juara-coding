package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mini-hris/config"
	"mini-hris/models"
)

type positionRequest struct {
	Title      string  `json:"title" binding:"required"`
	BaseSalary float64 `json:"base_salary" binding:"required,gt=0"`
}

func CreatePosition(ctx *gin.Context) {
	var req positionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.Title = normalizeText(req.Title)
	if req.Title == "" {
		respondError(ctx, http.StatusBadRequest, "title is required")
		return
	}

	position := models.Position{
		Title:      req.Title,
		BaseSalary: req.BaseSalary,
	}

	if err := config.DB.Create(&position).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "position created",
		"data":    position,
	})
}

func ListPositions(ctx *gin.Context) {
	var positions []models.Position
	if err := config.DB.Order("id ASC").Find(&positions).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": positions})
}

func UpdatePosition(ctx *gin.Context) {
	id, ok := parseUintParam(ctx, "id")
	if !ok {
		return
	}

	var position models.Position
	if err := config.DB.First(&position, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(ctx, http.StatusNotFound, "position not found")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var req positionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.Title = normalizeText(req.Title)
	if req.Title == "" {
		respondError(ctx, http.StatusBadRequest, "title is required")
		return
	}

	position.Title = req.Title
	position.BaseSalary = req.BaseSalary

	if err := config.DB.Save(&position).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "position updated",
		"data":    position,
	})
}

func DeletePosition(ctx *gin.Context) {
	id, ok := parseUintParam(ctx, "id")
	if !ok {
		return
	}

	var position models.Position
	if err := config.DB.First(&position, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(ctx, http.StatusNotFound, "position not found")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := config.DB.Delete(&position).Error; err != nil {
		if isForeignKeyConstraintError(err) {
			respondError(ctx, http.StatusConflict, "position is still used by employees")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "position deleted"})
}

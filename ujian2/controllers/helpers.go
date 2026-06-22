package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mini-hris/config"
	"mini-hris/models"
)

func respondError(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{"error": message})
}

func parseUintParam(ctx *gin.Context, name string) (uint, bool) {
	value := ctx.Param(name)
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		respondError(ctx, http.StatusBadRequest, "invalid "+name)
		return 0, false
	}

	return uint(id), true
}

func isUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(strings.ToLower(err.Error()), "unique constraint")
}

func isForeignKeyConstraintError(err error) bool {
	if err == nil {
		return false
	}

	return strings.Contains(strings.ToLower(err.Error()), "foreign key constraint")
}

func parseDate(value string) (time.Time, error) {
	return time.Parse("2006-01-02", value)
}

func parseClock(value string) (time.Time, error) {
	return time.Parse("15:04", value)
}

func parsePeriod(value string) (time.Time, error) {
	return time.Parse("2006-01", value)
}

func ensureEmployeeExists(employeeID uint) error {
	var employee models.Employee
	if err := config.DB.First(&employee, employeeID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("employee not found")
		}
		return err
	}

	return nil
}

func ensureDepartmentExists(departmentID uint) error {
	var department models.Department
	if err := config.DB.First(&department, departmentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("department not found")
		}
		return err
	}

	return nil
}

func ensurePositionExists(positionID uint) error {
	var position models.Position
	if err := config.DB.First(&position, positionID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("position not found")
		}
		return err
	}

	return nil
}

func normalizeText(value string) string {
	return strings.TrimSpace(value)
}

package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"mini-hris/config"
	"mini-hris/models"
)

const (
	presentAllowance = 50000.0
	lateDeduction    = 20000.0
	absentDeduction  = 100000.0
)

type salaryCalculationRequest struct {
	Period     string `json:"period" binding:"required"`
	EmployeeID *uint  `json:"employee_id" binding:"omitempty,gt=0"`
}

func CalculateSalaries(ctx *gin.Context) {
	var req salaryCalculationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	periodStart, err := parsePeriod(req.Period)
	if err != nil {
		respondError(ctx, http.StatusBadRequest, "period must use YYYY-MM format")
		return
	}

	var salaries []models.Salary
	err = config.DB.Transaction(func(tx *gorm.DB) error {
		query := tx.Preload("Position").Where("status = ?", "ACTIVE").Order("id ASC")
		if req.EmployeeID != nil {
			query = query.Where("id = ?", *req.EmployeeID)
		}

		var employees []models.Employee
		if err := query.Find(&employees).Error; err != nil {
			return err
		}

		if len(employees) == 0 {
			return gorm.ErrRecordNotFound
		}

		for _, employee := range employees {
			allowance, deductions, err := calculateAttendanceSummary(tx, employee.ID, req.Period)
			if err != nil {
				return err
			}

			approvedLeaveDays, err := calculateApprovedLeaveDays(tx, employee.ID, periodStart, req.Period)
			if err != nil {
				return err
			}

			salary := models.Salary{
				EmployeeID:  employee.ID,
				Period:      req.Period,
				BasicSalary: employee.Position.BaseSalary,
				Allowance:   allowance,
				Deductions:  deductions,
				NetSalary:   employee.Position.BaseSalary + allowance - deductions,
			}

			var existing models.Salary
			err = tx.Where("employee_id = ? AND period = ?", employee.ID, req.Period).First(&existing).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if errors.Is(err, gorm.ErrRecordNotFound) {
				if approvedLeaveDays > 0 {
					newBalance := employee.LeaveBalance - approvedLeaveDays
					if newBalance < 0 {
						newBalance = 0
					}

					if err := tx.Model(&models.Employee{}).
						Where("id = ?", employee.ID).
						Update("leave_balance", newBalance).Error; err != nil {
						return err
					}
				}

				if err := tx.Create(&salary).Error; err != nil {
					return err
				}
				continue
			}

			existing.BasicSalary = salary.BasicSalary
			existing.Allowance = salary.Allowance
			existing.Deductions = salary.Deductions
			existing.NetSalary = salary.NetSalary

			if err := tx.Save(&existing).Error; err != nil {
				return err
			}
		}

		salaryQuery := tx.Preload("Employee.Department").Preload("Employee.Position").
			Where("period = ?", req.Period).Order("employee_id ASC")
		if req.EmployeeID != nil {
			salaryQuery = salaryQuery.Where("employee_id = ?", *req.EmployeeID)
		}

		return salaryQuery.Find(&salaries).Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(ctx, http.StatusNotFound, "no active employees found for payroll calculation")
			return
		}

		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "salary calculation completed",
		"data":    salaries,
	})
}

func ListSalariesByPeriod(ctx *gin.Context) {
	period := ctx.Param("period")
	if _, err := parsePeriod(period); err != nil {
		respondError(ctx, http.StatusBadRequest, "period must use YYYY-MM format")
		return
	}

	var salaries []models.Salary
	if err := config.DB.Preload("Employee.Department").Preload("Employee.Position").
		Where("period = ?", period).Order("employee_id ASC").Find(&salaries).Error; err != nil {
		respondError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": salaries})
}

func calculateAttendanceSummary(tx *gorm.DB, employeeID uint, period string) (float64, float64, error) {
	var attendances []models.Attendance
	if err := tx.Where("employee_id = ? AND date LIKE ?", employeeID, period+"-%").
		Find(&attendances).Error; err != nil {
		return 0, 0, err
	}

	var allowance float64
	var deductions float64
	for _, attendance := range attendances {
		switch attendance.Status {
		case "PRESENT":
			allowance += presentAllowance
		case "LATE":
			deductions += lateDeduction
		case "ABSENT":
			deductions += absentDeduction
		}
	}

	return allowance, deductions, nil
}

func calculateApprovedLeaveDays(tx *gorm.DB, employeeID uint, periodStart time.Time, period string) (int, error) {
	var leaves []models.Leave
	if err := tx.Where("employee_id = ? AND status = ? AND start_date <= ? AND end_date >= ?",
		employeeID,
		"APPROVED",
		period+"-31",
		period+"-01",
	).Find(&leaves).Error; err != nil {
		return 0, err
	}

	periodEnd := periodStart.AddDate(0, 1, -1)
	totalDays := 0
	for _, leave := range leaves {
		startDate, err := parseDate(leave.StartDate)
		if err != nil {
			return 0, err
		}

		endDate, err := parseDate(leave.EndDate)
		if err != nil {
			return 0, err
		}

		if startDate.Before(periodStart) {
			startDate = periodStart
		}

		if endDate.After(periodEnd) {
			endDate = periodEnd
		}

		if endDate.Before(startDate) {
			continue
		}

		totalDays += int(endDate.Sub(startDate).Hours()/24) + 1
	}

	return totalDays, nil
}

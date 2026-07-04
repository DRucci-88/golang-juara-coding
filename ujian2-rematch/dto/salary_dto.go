package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type SalaryCalculateRequest struct {
	Period string `json:"period" binding:"required,payroll_period"`
}

func SalaryPayrollPeriodValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	// return regexp.MustCompile(`^\d{4}-(0[1-9]|1[0-2])$`).MatchString(value)
	_, err := time.Parse("2006-01", value)
	return err == nil
}

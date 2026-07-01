package model

import (
	"time"

	"gorm.io/gorm"
)

type Salary struct {
	gorm.Model

	EmployeeID uint
	Employee   Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Period      time.Time `gorm:"type:varchar(10);not null"`
	BasicSalary float64   `gorm:"type:numeric(12,2);not null"`
	Allowance   float64   `gorm:"type:numeric(12,2);not null"`
	Deductions  float64   `gorm:"type:numeric(12,2);not null"`
	NetSalary   float64   `gorm:"type:numeric(12,2);not null"`
}

type SalaryPreload string

const (
	SalaryPreloadEmployee SalaryPreload = "Employee"
)

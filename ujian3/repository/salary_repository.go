package repository

import (
	"time"

	"gorm.io/gorm"
)

type SalaryDB struct {
	ID uint `gorm:"primaryKey"`

	EmployeeID uint       `gorm:"not null;uniqueIndex:idx_salary_employee_period"`
	Employee   EmployeeDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Period      string  `gorm:"not null;size:7;uniqueIndex:idx_salary_employee_period"`
	BasicSalary float64 `gorm:"not null"`
	Allowance   float64 `gorm:"not null"`
	Deductions  float64 `gorm:"not null"`
	NetSalary   float64 `gorm:"not null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (SalaryDB) TableName() string { return "salaries" }

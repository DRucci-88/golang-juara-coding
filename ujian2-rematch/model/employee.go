package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	NIK      string         `gorm:"type:varchar(50);not null"`
	FullName string         `gorm:"type:varchar(100);not null"`
	Email    string         `gorm:"type:varchar(100);not null"`
	Status   EmployeeStatus `gorm:"type:varchar(20);default:'ACTIVE';not null"`

	DepartmentID uint
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	PositionID uint
	Position   Position `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Attendances []Attendance `gorm:"foreignKey:EmployeeID"`
	Leaves      []Leave      `gorm:"foreignKey:EmployeeID"`
	Salaries    []Salary     `gorm:"foreignKey:EmployeeID"`
}

type EmployeeStatus string

const (
	EmployeeStatusActive     EmployeeStatus = "ACTIVE"
	EmployeeStatusSuspended  EmployeeStatus = "SUSPENDED"
	EmployeeStatusTerminated EmployeeStatus = "TERMINATED"
)

package model

import (
	"errors"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	NIK      string         `gorm:"type:varchar(50);not null"`
	FullName string         `gorm:"type:varchar(100);not null"`
	Status   EmployeeStatus `gorm:"type:varchar(20);default:'ACTIVE';not null"`

	UserID uint
	User   *User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	DepartmentID uint
	Department   *Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	PositionID uint
	Position   *Position `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

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

type EmployeePreload string

const (
	EmployeePreloadDepartment EmployeePreload = "Department"
	EmployeePreloadPosition   EmployeePreload = "Position"

	EmployeePreloadAttendance EmployeePreload = "Attendances"
	EmployeePreloadLeave      EmployeePreload = "Leaves"
	EmployeePreloadSalary     EmployeePreload = "Salaries"
)

var (
	ErrEmployeeNotFound = errors.New("Employee Not Found")
)

package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name string `gorm:"type:varchar(50);not null"`
	Code string `gorm:"type:varchar(50);not null;uniqueIndex"`

	Employees []Employee `gorm:"foreignKey:DepartmentID"`
}

type DepartmentPreload string

const (
	DepartmentPreloadEmployees DepartmentPreload = "Employees"
)

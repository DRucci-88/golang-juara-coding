package repository

import (
	"time"

	"gorm.io/gorm"
)

type DepartmentDB struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Code string `gorm:"not null;uniqueIndex"`

	Employees []EmployeeDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (DepartmentDB) TableName() string { return "departments" }

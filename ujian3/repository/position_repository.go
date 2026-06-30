package repository

import (
	"time"

	"gorm.io/gorm"
)

type PositionDB struct {
	ID         uint    `gorm:"primaryKey"`
	Title      string  `gorm:"not null"`
	BaseSalary float64 `gorm:"not null"`

	Employees []EmployeeDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (PositionDB) TableName() string { return "position" }

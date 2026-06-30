package repository

import (
	"time"

	"gorm.io/gorm"
)

type LeaveDB struct {
	ID uint `gorm:"primaryKey"`

	EmployeeID uint       `gorm:"not null"`
	Employee   EmployeeDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	StartDate string `gorm:"not null;size:10"`
	EndDate   string `gorm:"not null;size:10"`
	Reason    string `gorm:"not null"`
	Status    string `gorm:"not null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (LeaveDB) TableName() string { return "leaves" }

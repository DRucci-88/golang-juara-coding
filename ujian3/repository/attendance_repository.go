package repository

import (
	"time"

	"gorm.io/gorm"
)

type AttendanceDB struct {
	ID uint `gorm:"primaryKey"`

	EmployeeID uint       `gorm:"not null;uniqueIndex:idx_attendance_employee_date"`
	Employee   EmployeeDB `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Date     string `gorm:"not null;size:10;uniqueIndex:idx_attendance_employee_date"`
	CheckIn  string `gorm:"not null;size:5"`
	CheckOut string `gorm:"not null;size:5"`
	Status   string `gorm:"not null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (AttendanceDB) TableName() string { return "attendances" }

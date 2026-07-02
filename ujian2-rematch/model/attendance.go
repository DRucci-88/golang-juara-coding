package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model

	EmployeeID uint
	Employee   *Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Date     time.Time        `gorm:"type:date;not null"`
	CheckIn  sql.NullTime     `gorm:"type:time"`
	CheckOut sql.NullTime     `gorm:"type:time"`
	Status   AttendanceStatus `gorm:"type:varchar(10);not null"`
}

type AttendanceStatus string

const (
	AttendanceStatusPresent AttendanceStatus = "PRESENT"
	AttendanceStatusLate    AttendanceStatus = "LATE"
	AttendanceStatusAbsent  AttendanceStatus = "ABSENT"
)

type AttendancePreload string

const (
	AttendancePreloadEmployee AttendancePreload = "Employee"
)

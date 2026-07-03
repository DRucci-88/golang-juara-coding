package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	gorm.Model

	EmployeeID uint
	Employee   *Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCAFE,OnDelete:RESTRICT"`

	StartDate time.Time   `gorm:"type:date;not null"`
	EndDate   time.Time   `gorm:"type:date;not null"`
	Reason    string      `gorm:"type:varchar(255)"`
	Status    LeaveStatus `gorm:"type:varchar(20);default:'PENDING'"`
}

type LeaveStatus string

const (
	LeaveStatusPending  LeaveStatus = "PENDING"
	LeaveStatusApproved LeaveStatus = "APPROVED"
	LeaveStatusRejected LeaveStatus = "REJECTED"
)

type LeavePreload string

const (
	LeavePreloadEmployee LeavePreload = "Employee"
)

var (
	ErrLeaveNotFound    = errors.New("Leave Not Found")
	ErrLeaveOverlapping = errors.New("Leave Overlapping")
)

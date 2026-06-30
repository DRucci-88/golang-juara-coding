package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	gorm.Model

	EmployeeID uint
	Employee   Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCAFE,OnDelete:RESTRICT"`

	StartDate time.Time      `gorm:"type:date;not null"`
	EndDate   time.Time      `gorm:"type:date;not null"`
	Reason    sql.NullString `gorm:"type:varchar(255)"`
	Status    LeaveStatus    `gorm:"type:varchar(20);default:'PENDING'"`
}

type LeaveStatus string

const (
	LeaveStatusPending  LeaveStatus = "PENDING"
	LeaveStatusApproved LeaveStatus = "APPROVED"
	LeaveStatusRejected LeaveStatus = "REJECTED"
)

type LeaveRepository[T any] interface {

	// INSERT INTO @@table
	Create(leave *Leave) error

	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	FindByID(id uint) (*Leave, error)

	// SELECT * FROM @@table
	FindAll() ([]Leave, error)

	// SELECT * FROM @@table WHERE employee_id = @employeeID
	FindByEmployeeID(employeeID uint) ([]Leave, error)

	// UPDATE @@table
	Update(leave *Leave) error

	// DELETE FROM @@table WHERE id = @id
	Delete(id uint) error
}
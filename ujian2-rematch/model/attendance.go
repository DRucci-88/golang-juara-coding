package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model

	EmployeeID uint
	Employee   Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Date     time.Time        `gorm:"type:date;not null"`
	CheckIn  sql.NullString   `gorm:"varchar(10)"`
	CheckOut sql.NullString   `gorm:"varchar(10)"`
	Status   AttendanceStatus `gorm:"varchar(10);not null"`
}

type AttendanceStatus string

const (
	AttendanceStatusPresent AttendanceStatus = "PRESENT"
	AttendanceStatusLate    AttendanceStatus = "LATE"
	AttendanceStatusAbsent  AttendanceStatus = "ABSENT"
)

type AttendanceRepository[T Attendance] interface {

	// INSERT INTO @@table
	Create(tx *gorm.DB, attendance *Attendance) error

	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	FindByID(id uint) (*Attendance, error)

	// SELECT * FROM @@table
	FindAll() ([]Attendance, error)

	// SELECT * FROM @@table WHERE employee_id = @employeeID
	FindByEmployeeID(employeeID uint) ([]Attendance, error)

	// SELECT * FROM @@table WHERE employee_id = @employeeID AND date = @date LIMIT 1
	FindByEmployeeAndDate(employeeID uint, date time.Time) (*Attendance, error)

	// UPDATE @@table
	Update(attendance *Attendance) error

	// DELETE FROM @@table WHERE id = @id
	Delete(id uint) error
}

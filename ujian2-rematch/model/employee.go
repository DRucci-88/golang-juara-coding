package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model

	NIK      string         `gorm:"type:varchar(50);not null"`
	FullName string         `gorm:"type:varchar(100);not null"`
	Email    string         `gorm:"type:varchar(100);not null"`
	Status   EmployeeStatus `gorm:"type:varchar(20);default:'ACTIVE';not null"`

	DepartmentID uint
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	PositionID uint
	Position   Position `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

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

type EmployeeRepository[T any] interface {

	// INSERT INTO @@table
	Create(employee *Employee) error

	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	FindByID(id uint) (*Employee, error)

	// SELECT * FROM @@table
	FindAll() ([]Employee, error)

	// SELECT * FROM @@table WHERE email = @email LIMIT 1
	FindByEmail(email string) (*Employee, error)

	// SELECT * FROM @@table WHERE nik = @nik LIMIT 1
	FindByNIK(nik string) (*Employee, error)

	// SELECT * FROM @@table WHERE department_id = @departmentID
	FindByDepartmentID(departmentID uint) ([]Employee, error)

	// SELECT * FROM @@table WHERE position_id = @positionID
	FindByPositionID(positionID uint) ([]Employee, error)

	// UPDATE @@table
	Update(employee *Employee) error

	// DELETE FROM @@table WHERE id = @id
	Delete(id uint) error
}

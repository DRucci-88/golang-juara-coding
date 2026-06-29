package domain

import (
	"errors"
	"time"
)

var (
	ErrEmployeeNIKDuplicate = errors.New("NIK Duplicated")
)

type Employee struct {
	ID uint

	UserID uint
	User   *User

	NIK      string
	FullName string
	Status   string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type EmployeeRepository interface {
	Create(employee *Employee) (*Employee, error)
	FindByNIK(nik string) (*Employee, error)
	// FindByID(id uint) (*Employee, error)
	// FindByUserID(userID uint) (*Employee, error)
	// Update(employee *Employee) error
	// Delete(id uint) error
}

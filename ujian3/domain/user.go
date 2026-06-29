package domain

import (
	"errors"
	"time"
)

var (
	ErrUserEmailDuplicate = errors.New("Email Duplicate")
	ErrUserNotFound       = errors.New("User Not Found")
)

type Role = string

const (
	RoleHRD      Role = "HRD"
	RoleEMPLOYEE Role = "EMPLOYEE"
)

type User struct {
	ID       uint
	Email    string
	Password string
	Role     string

	Employee *Employee

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserRepository interface {
	Create(user *User, hashPassword string) (*User, error)
	// FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	// Update(user *User) error
	// Delete(id uint) error
}

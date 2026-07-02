package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role     UserRole `gorm:"type:varchar(20);not null"`
	Email    string   `gorm:"type:varchar(100);not null"`
	Password string   `gorm:"type:varchar(255)"`

	Employee *Employee `gorm:"foreignKey:UserID"`
}

type UserRole string

const (
	UserRoleAdmin    UserRole = "ADMIN"
	UserRoleEmployee UserRole = "EMPLOYEE"
)

type UserPreload string

const (
	UserPreloadEmployee UserPreload = "Employee"
)

var (
	ErrUserNotFound = errors.New("User Not Found")
)

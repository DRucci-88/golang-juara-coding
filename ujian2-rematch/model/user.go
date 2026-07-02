package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Role     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(255)"`

	Employee *Employee `gorm:"fo"`
}

type UserRole string

const (
	UserRoleAdmin    UserRole = "ADMIN"
	UserRoleEmployee UserRole = "EMPLOYEE"
)

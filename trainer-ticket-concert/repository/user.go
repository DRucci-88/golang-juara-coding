package repository

import (
	"trainer-ticket-concert/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

package repository

import (
	"praktikum/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(tx *gorm.DB, user *model.User) error
	FindById(tx *gorm.DB, id uint) (*model.User, error)
	FindByEmail(tx *gorm.DB, email string) (*model.User, error)
}

type userRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) Create(tx *gorm.DB, user *model.User) error {
	return tx.Create(user).Error
}

func (r *userRepositoryImpl) FindById(tx *gorm.DB, id uint) (*model.User, error) {
	var user model.User
	err := tx.First(&user, id).Error
	return &user, err
}

func (r *userRepositoryImpl) FindByEmail(tx *gorm.DB, email string) (*model.User, error) {
	var user model.User
	err := tx.Where("email = ?", email).First(&user).Error
	return &user, err
}

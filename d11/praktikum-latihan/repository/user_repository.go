package repository

import (
	"praktikum/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(tx *gorm.DB, id uint) (*model.User, error)
}

type userRepositoryImpl struct {
}

func NewUserReposutory() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) FindById(tx *gorm.DB, id uint) (*model.User, error) {
	var user model.User
	err := tx.First(&user, id).Error
	return &user, err
}

package repository

import (
	"praktikum/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(tx *gorm.DB, order *model.Order) error
	FindByID(tx *gorm.DB, id uint) (*model.Order, error)
	Update(tx *gorm.DB, order *model.Order) error
}

type orderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &orderRepositoryImpl{}
}

func (r *orderRepositoryImpl) Create(tx *gorm.DB, order *model.Order) error {
	return tx.Create(order).Error
}

func (r *orderRepositoryImpl) FindByID(tx *gorm.DB, id uint) (*model.Order, error) {
	var order model.Order
	err := tx.First(&order, id).Error
	return &order, err
}

func (r *orderRepositoryImpl) Update(tx *gorm.DB, order *model.Order) error {
	return tx.Save(order).Error
}

package repository

import (
	"praktikum/model"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(tx *gorm.DB, orderItem *model.OrderItem) error
}

type orderItemRepositoryImpl struct {
}

func NewOrderItemRepository() OrderItemRepository {
	return &orderItemRepositoryImpl{}
}

func (r *orderItemRepositoryImpl) Create(tx *gorm.DB, orderItem *model.OrderItem) error {
	return tx.Create(orderItem).Error
}

package repository

import (
	"praktikum/dto"
	"praktikum/model"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(tx *gorm.DB, orderItem *model.OrderItem) error
	AnalyticsPopularProduct(tx *gorm.DB, productID uint) (*[]dto.ProductSales, error)
}

type orderItemRepositoryImpl struct {
}

func NewOrderItemRepository() OrderItemRepository {
	return &orderItemRepositoryImpl{}
}

func (r *orderItemRepositoryImpl) Create(tx *gorm.DB, orderItem *model.OrderItem) error {
	return tx.Create(orderItem).Error
}

func (r *orderItemRepositoryImpl) AnalyticsPopularProduct(tx *gorm.DB, productID uint) (*[]dto.ProductSales, error) {
	var result []dto.ProductSales

	err := tx.
		Model(&model.OrderItem{}).
		Select("product_id, SUM(quantity) as total_quantity").
		Group("product_id").
		Scan(&result).
		Error

	return &result, err

}

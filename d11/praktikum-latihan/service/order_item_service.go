package service

import (
	"praktikum/dto"
	"praktikum/repository"

	"gorm.io/gorm"
)

type OrderItemService interface {
	AnalyticsPopularProduct(productID int) (*[]dto.ProductSales, error)
}

type orderItemServiceImpl struct {
	db            *gorm.DB
	orderItemRepo repository.OrderItemRepository
}

func NewOrderItemService(
	db *gorm.DB,
	orderItemRepo repository.OrderItemRepository,
) OrderItemService {
	return &orderItemServiceImpl{
		db:            db,
		orderItemRepo: orderItemRepo,
	}
}

func (s *orderItemServiceImpl) AnalyticsPopularProduct(productID int) (*[]dto.ProductSales, error) {
	return s.orderItemRepo.AnalyticsPopularProduct(s.db, uint(productID))
}

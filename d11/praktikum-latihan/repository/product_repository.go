package repository

import (
	"praktikum/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByID(tx *gorm.DB, id uint) (*model.Product, error)
	Update(tx *gorm.DB, product *model.Product) error
}

type productRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &productRepositoryImpl{}
}

func (r *productRepositoryImpl) FindByID(tx *gorm.DB, id uint) (*model.Product, error) {
	var product *model.Product
	err := tx.First(product, id).Error

	return product, err
}

func (r *productRepositoryImpl) Update(tx *gorm.DB, product *model.Product) error {
	return tx.Save(product).Error
}

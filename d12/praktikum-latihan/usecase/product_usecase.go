package usecase

import (
	"context"
	"praktikum/domain"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(repo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{repo: repo}
}
func (u *productUsecase) Create(ctx context.Context, sku string, name string, price float64, stock int) (*domain.Product, error) {
	// Menjalankan logika bisnis tambahan (misal validasi bisnis)
	newProduct := &domain.Product{
		SKU:   sku,
		Name:  name,
		Price: price,
		Stock: stock,
	}
	err := u.repo.Create(ctx, newProduct)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (u *productUsecase) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	return u.repo.GetByID(ctx, id)
}

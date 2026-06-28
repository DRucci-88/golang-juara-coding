package repository

import (
	"context"
	"errors"
	"praktikum/domain"
	"time"

	"gorm.io/gorm"
)

// GORM DB Model (Hanya ada di layer repository)
type ProductDB struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	SKU       string         `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name      string         `gorm:"type:varchar(100);not null"`
	Price     float64        `gorm:"type:numeric(12,2);not null"`
	Stock     int            `gorm:"type:int;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Mengatur nama tabel database GORM
func (ProductDB) TableName() string { return "products" }

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, p *domain.Product) error {
	dbModel := ProductDB{
		SKU:   p.SKU,
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	}
	// Cek duplikasi SKU
	var count int64
	r.db.WithContext(ctx).Model(&ProductDB{}).Where("sku = ?",
		p.SKU).Count(&count)
	if count > 0 {
		return domain.ErrSKUDuplicate
	}
	err := r.db.WithContext(ctx).Create(&dbModel).Error
	if err != nil {
		return err
	}
	// Salin ID dan waktu yang digenerate oleh DB ke Entity Domain
	p.ID = dbModel.ID
	p.CreatedAt = dbModel.CreatedAt
	p.UpdatedAt = dbModel.UpdatedAt
	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	var dbModel ProductDB
	err := r.db.WithContext(ctx).First(&dbModel, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &domain.Product{
		ID:        dbModel.ID,
		SKU:       dbModel.SKU,
		Name:      dbModel.Name,
		Price:     dbModel.Price,
		Stock:     dbModel.Stock,
		CreatedAt: dbModel.CreatedAt,
		UpdatedAt: dbModel.UpdatedAt,
	}, nil
}

func (r *productRepository) GetBySKU(ctx context.Context, sku string) (*domain.Product, error) {
	var dbModel ProductDB
	err := r.db.WithContext(ctx).Where("sku = ?", sku).First(&dbModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}
	return &domain.Product{
		ID:        dbModel.ID,
		SKU:       dbModel.SKU,
		Name:      dbModel.Name,
		Price:     dbModel.Price,
		Stock:     dbModel.Stock,
		CreatedAt: dbModel.CreatedAt,
		UpdatedAt: dbModel.UpdatedAt,
	}, nil
}

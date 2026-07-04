package repository

import (
	"context"
	"ujian2_rematch/model"

	"gorm.io/gorm"
)

type SalaryRepository struct {
	db    *gorm.DB
	query gorm.Interface[model.Salary]
}

func (r *RepositoryManager) Salary() *SalaryRepository {
	return &SalaryRepository{
		db:    r.db,
		query: gorm.G[model.Salary](r.db),
	}
}

func (r *SalaryRepository) queryWithPreloads(
	preloads ...model.SalaryPreload,
) gorm.ChainInterface[model.Salary] {
	var chain gorm.ChainInterface[model.Salary] = r.query.Scopes()
	for _, preload := range preloads {
		chain = chain.Preload(string(preload), nil)
	}
	return chain
}

func (r *SalaryRepository) Create(
	ctx context.Context,
	salary *model.Salary,
) error {
	return r.query.Create(ctx, salary)
}

func (r *SalaryRepository) CreateInBatches(
	ctx context.Context,
	batchSize int,
	salaries []model.Salary,
) error {
	return r.query.CreateInBatches(ctx, &salaries, batchSize)
}

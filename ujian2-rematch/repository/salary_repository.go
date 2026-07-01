package repository

import (
	"context"
	"time"

	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type SalaryRepository struct {
	db *gorm.DB
}

func (r *RepositoryManager) Salary() *SalaryRepository {
	return &SalaryRepository{
		db: r.db,
	}
}

func (r *SalaryRepository) preload(
	db *gorm.DB,
	preloads ...model.SalaryPreload,
) *gorm.DB {

	for _, preload := range preloads {
		db = db.Preload(string(preload))
	}

	return db
}

func (r *SalaryRepository) Create(
	ctx context.Context,
	salary *model.Salary,
) error {

	return gorm.G[model.Salary](r.db).
		Create(ctx, salary)
}

func (r *SalaryRepository) Update(
	ctx context.Context,
	salary *model.Salary,
) error {

	_, err := gorm.G[model.Salary](r.db).
		Where(generated.Salary.ID.Eq(salary.ID)).
		Updates(ctx, *salary)

	return err
}

func (r *SalaryRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	_, err := gorm.G[model.Salary](r.db).
		Where(generated.Salary.ID.Eq(id)).
		Delete(ctx)

	return err
}

func (r *SalaryRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.SalaryPreload,
) (*model.Salary, error) {

	salary, err := gorm.G[model.Salary](
		r.preload(r.db, preloads...),
	).
		Where(generated.Salary.ID.Eq(id)).
		First(ctx)

	return &salary, err
}

func (r *SalaryRepository) FindAll(
	ctx context.Context,
	preloads ...model.SalaryPreload,
) ([]model.Salary, error) {

	return gorm.G[model.Salary](
		r.preload(r.db, preloads...),
	).
		Find(ctx)
}

func (r *SalaryRepository) FindAllByEmployeeID(
	ctx context.Context,
	employeeID uint,
	preloads ...model.SalaryPreload,
) ([]model.Salary, error) {

	return gorm.G[model.Salary](
		r.preload(r.db, preloads...),
	).
		Where(generated.Salary.EmployeeID.Eq(employeeID)).
		Find(ctx)
}

func (r *SalaryRepository) FindByEmployeeAndPeriod(
	ctx context.Context,
	employeeID uint,
	period time.Time,
	preloads ...model.SalaryPreload,
) (*model.Salary, error) {

	salary, err := gorm.G[model.Salary](
		r.preload(r.db, preloads...),
	).
		Where(generated.Salary.EmployeeID.Eq(employeeID)).
		Where(generated.Salary.Period.Eq(period)).
		First(ctx)

	return &salary, err
}

func (r *SalaryRepository) FindAllByPeriod(
	ctx context.Context,
	period time.Time,
	preloads ...model.SalaryPreload,
) ([]model.Salary, error) {

	return gorm.G[model.Salary](
		r.preload(r.db, preloads...),
	).
		Where(generated.Salary.Period.Eq(period)).
		Find(ctx)
}

func (r *SalaryRepository) FindAllByYear(
	ctx context.Context,
	year int,
	preloads ...model.SalaryPreload,
) ([]model.Salary, error) {

	start := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(1, 0, 0)

	return gorm.G[model.Salary](
		r.preload(r.db, preloads...),
	).
		Where(generated.Salary.Period.Gte(start)).
		Where(generated.Salary.Period.Lt(end)).
		Find(ctx)
}

func (r *SalaryRepository) ExistsPayroll(
	ctx context.Context,
	employeeID uint,
	period time.Time,
) (bool, error) {

	count, err := gorm.G[model.Salary](r.db).
		Where(generated.Salary.EmployeeID.Eq(employeeID)).
		Where(generated.Salary.Period.Eq(period)).
		Count(ctx, "*")

	return count > 0, err
}

func (r *SalaryRepository) Count(
	ctx context.Context,
) (int64, error) {

	return gorm.G[model.Salary](r.db).
		Count(ctx, "*")
}

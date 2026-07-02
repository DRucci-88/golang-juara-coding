package repository

import (
	"context"
	"errors"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db    *gorm.DB
	query gorm.Interface[model.Department]
}

func (r *RepositoryManager) Department() *DepartmentRepository {
	return &DepartmentRepository{
		db:    r.db,
		query: gorm.G[model.Department](r.db),
	}
}

func (r *DepartmentRepository) queryWithPreloads(
	preloads ...model.DepartmentPreload,
) gorm.ChainInterface[model.Department] {
	var chain gorm.ChainInterface[model.Department] = r.query.Scopes()
	for _, preload := range preloads {
		chain = chain.Preload(string(preload), nil)
	}
	return chain
}

func (r *DepartmentRepository) Create(
	ctx context.Context,
	department *model.Department,
) error {
	err := gorm.G[model.Department](r.db).
		Create(ctx, department)
	return err
}

func (r *DepartmentRepository) Update(
	ctx context.Context,
	id uint,
	department *model.Department,
) (int, error) {
	rows, err := gorm.G[model.Department](r.db).
		Where(generated.Department.ID.Eq(id)).
		Updates(ctx, *department)

	if rows == 0 {
		return rows, model.ErrDepartmentNotFound
	}

	return rows, err
}

func (r *DepartmentRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	_, err := gorm.G[model.Department](r.db).
		Where(generated.Department.ID.Eq(id)).
		Delete(ctx)

	return err
}

func (r *DepartmentRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.DepartmentPreload,
) (*model.Department, error) {

	department, err := r.queryWithPreloads(preloads...).
		Where(generated.Department.ID.Eq(id)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrDepartmentNotFound
	}

	return &department, err
}

func (r *DepartmentRepository) FindAll(
	ctx context.Context,
	preloads ...model.DepartmentPreload,
) ([]model.Department, error) {

	return r.queryWithPreloads(preloads...).
		Find(ctx)
}

func (r *DepartmentRepository) FindByCode(
	ctx context.Context,
	code string,
	preloads ...model.DepartmentPreload,
) (*model.Department, error) {

	department, err := r.queryWithPreloads(preloads...).
		Where(generated.Department.Code.Eq(code)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrPositionNotFound
	}

	return &department, err
}

func (r *DepartmentRepository) FindByName(
	ctx context.Context,
	name string,
	preloads ...model.DepartmentPreload,
) (*model.Department, error) {

	department, err := r.queryWithPreloads(preloads...).
		Where(generated.Department.Name.Eq(name)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrPositionNotFound
	}

	return &department, err
}

func (r *DepartmentRepository) ExistsByCode(
	ctx context.Context,
	code string,
) (bool, error) {

	count, err := gorm.G[model.Department](r.db).
		Where(generated.Department.Code.Eq(code)).
		Count(ctx, "*")

	return count > 0, err
}

func (r *DepartmentRepository) ExistsByName(
	ctx context.Context,
	name string,
) (bool, error) {

	count, err := gorm.G[model.Department](r.db).
		Where(generated.Department.Name.Eq(name)).
		Count(ctx, "*")

	return count > 0, err
}

func (r *DepartmentRepository) Count(
	ctx context.Context,
) (int64, error) {

	return gorm.G[model.Department](r.db).
		Count(ctx, "*")
}

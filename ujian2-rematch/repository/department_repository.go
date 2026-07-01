package repository

import (
	"context"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func (r *RepositoryManager) Department() *DepartmentRepository {
	return &DepartmentRepository{
		db: r.db,
	}
}

func (r *DepartmentRepository) preload(
	db *gorm.DB,
	preloads ...model.DepartmentPreload,
) *gorm.DB {

	for _, preload := range preloads {
		db = db.Preload(string(preload))
	}

	return db
}

func (r *DepartmentRepository) Create(
	ctx context.Context,
	department *model.Department,
) error {

	return gorm.G[model.Department](r.db).
		Create(ctx, department)
}

func (r *DepartmentRepository) Update(
	ctx context.Context,
	department *model.Department,
) error {

	_, err := gorm.G[model.Department](r.db).
		Where(generated.Department.ID.Eq(department.ID)).
		Updates(ctx, *department)

	return err
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

	department, err := gorm.G[model.Department](
		r.preload(r.db, preloads...),
	).
		Where(generated.Department.ID.Eq(id)).
		First(ctx)

	return &department, err
}

func (r *DepartmentRepository) FindAll(
	ctx context.Context,
	preloads ...model.DepartmentPreload,
) ([]model.Department, error) {

	return gorm.G[model.Department](
		r.preload(r.db, preloads...),
	).
		Find(ctx)
}

func (r *DepartmentRepository) FindByCode(
	ctx context.Context,
	code string,
	preloads ...model.DepartmentPreload,
) (*model.Department, error) {

	department, err := gorm.G[model.Department](
		r.preload(r.db, preloads...),
	).
		Where(generated.Department.Code.Eq(code)).
		First(ctx)

	return &department, err
}

func (r *DepartmentRepository) FindByName(
	ctx context.Context,
	name string,
	preloads ...model.DepartmentPreload,
) (*model.Department, error) {

	department, err := gorm.G[model.Department](
		r.preload(r.db, preloads...),
	).
		Where(generated.Department.Name.Eq(name)).
		First(ctx)

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

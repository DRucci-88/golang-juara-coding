package repository

import (
	"context"
	"errors"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

// Factory method from RepositoryManager
func (r *RepositoryManager) Employee() *EmployeeRepository {
	return &EmployeeRepository{
		db: r.db,
	}
}

func (r *EmployeeRepository) preload(
	db *gorm.DB,
	preloads ...model.EmployeePreload,
) *gorm.DB {

	for _, preload := range preloads {
		db = db.Preload(string(preload))
	}

	return db
}

// Create
func (r *EmployeeRepository) Create(
	ctx context.Context,
	employee *model.Employee,
) error {
	return gorm.G[model.Employee](r.db).
		Create(ctx, employee)
}

// Update
func (r *EmployeeRepository) Update(
	ctx context.Context,
	id uint,
	employee *model.Employee,
) (int, error) {

	rows, err := gorm.G[model.Employee](r.db).
		Where(generated.Employee.ID.Eq(id)).
		Updates(ctx, *employee)

	if rows == 0 {
		return rows, model.ErrEmployeeNotFound
	}

	return rows, err
}

// Delete
func (r *EmployeeRepository) Delete(
	ctx context.Context,
	id uint,
) (int, error) {
	return gorm.G[model.Employee](r.db).
		Where(generated.Employee.ID.Eq(id)).
		Delete(ctx)
}

// Find By ID
func (r *EmployeeRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.EmployeePreload,
) (*model.Employee, error) {

	employee, err := gorm.G[model.Employee](
		r.preload(r.db, preloads...),
	).
		Where(generated.Employee.ID.Eq(id)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrEmployeeNotFound
	}

	return &employee, err
}

// Find All
func (r *EmployeeRepository) FindAll(
	ctx context.Context,
	filter *dto.EmployeeFilterRequest,
	preloads ...model.EmployeePreload,
) ([]model.Employee, error) {

	db := r.preload(r.db, preloads...)

	if filter == nil {
		return gorm.G[model.Employee](db).Find(ctx)
	}

	if filter.Search != nil {
		db = db.Where(
			generated.Employee.FullName.Like("%" + *filter.Search + "%"),
		)
	}

	if filter.DepartmentID != nil {
		db = db.Where(
			generated.Employee.DepartmentID.Eq(*filter.DepartmentID),
		)
	}

	if filter.PositionID != nil {
		db = db.Where(
			generated.Employee.PositionID.Eq(*filter.PositionID),
		)
	}

	if filter.Status != nil {
		db = db.Where(
			generated.Employee.Status.Eq(string(*filter.Status)),
		)
	}

	return gorm.G[model.Employee](db).Find(ctx)

}

// Find By Email
func (r *EmployeeRepository) FindByEmail(
	ctx context.Context,
	email string,
	preloads ...model.EmployeePreload,
) (*model.Employee, error) {
	employee, err := gorm.G[model.Employee](
		r.preload(r.db, preloads...),
	).
		Where(generated.Employee.Email.Eq(email)).
		First(ctx)
	return &employee, err
}

// Find By NIK
func (r *EmployeeRepository) FindByNIK(
	ctx context.Context,
	nik string,
	preloads ...model.EmployeePreload,
) (*model.Employee, error) {
	employee, err := gorm.G[model.Employee](
		r.preload(r.db, preloads...),
	).
		Where(generated.Employee.NIK.Eq(nik)).
		First(ctx)
	return &employee, err
}

// Find By Department
func (r *EmployeeRepository) FindAllByDepartmentID(
	ctx context.Context,
	departmentID uint,
	preloads ...model.EmployeePreload,
) ([]model.Employee, error) {
	return gorm.G[model.Employee](
		r.preload(r.db, preloads...),
	).
		Where(generated.Employee.DepartmentID.Eq(departmentID)).
		Find(ctx)
}

// Find By Position
func (r *EmployeeRepository) FindAllByPositionID(
	ctx context.Context,
	positionID uint,
	preloads ...model.EmployeePreload,
) ([]model.Employee, error) {

	return gorm.G[model.Employee](
		r.preload(r.db, preloads...),
	).
		Where(generated.Employee.PositionID.Eq(positionID)).
		Find(ctx)
}

// Exists By Email
func (r *EmployeeRepository) ExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {
	count, err := gorm.G[model.Employee](r.db).
		Where(generated.Employee.Email.Eq(email)).
		Count(ctx, "*")
	return count > 0, err
}

// Exists By NIK
func (r *EmployeeRepository) ExistsByNIK(
	ctx context.Context,
	nik string,
) (bool, error) {
	count, err := gorm.G[model.Employee](r.db).
		Where(generated.Employee.NIK.Eq(nik)).
		Count(ctx, "*")
	return count > 0, err
}

// Count
func (r *EmployeeRepository) Count(
	ctx context.Context,
) (int64, error) {

	return gorm.G[model.Employee](r.db).
		Count(ctx, "*")
}

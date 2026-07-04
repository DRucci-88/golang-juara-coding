package repository

import (
	"context"
	"errors"
	"time"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeeRepository struct {
	db    *gorm.DB
	query gorm.Interface[model.Employee]
}

// Factory method from RepositoryManager
func (r *RepositoryManager) Employee() *EmployeeRepository {
	return &EmployeeRepository{
		db:    r.db,
		query: gorm.G[model.Employee](r.db),
	}
}

func (r *EmployeeRepository) queryWithPreloads(
	preloads ...model.EmployeePreload,
) gorm.ChainInterface[model.Employee] {
	var chain gorm.ChainInterface[model.Employee] = r.query.Scopes()
	for _, preload := range preloads {
		chain = chain.Preload(string(preload), nil)
	}
	return chain
}

// Create
func (r *EmployeeRepository) Create(
	ctx context.Context,
	employee *model.Employee,
) error {
	return r.query.
		Create(ctx, employee)
}

// Update
func (r *EmployeeRepository) Update(
	ctx context.Context,
	id uint,
	employee *model.Employee,
) (int, error) {

	rows, err := r.query.
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
	return r.query.
		Where(generated.Employee.ID.Eq(id)).
		Delete(ctx)
}

// Find By ID
func (r *EmployeeRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.EmployeePreload,
) (*model.Employee, error) {

	employee, err := r.queryWithPreloads(preloads...).
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

	chain := r.queryWithPreloads(preloads...)

	if filter == nil {
		return chain.Find(ctx)
	}

	if filter.Search != nil {
		chain = chain.Where(
			generated.Employee.FullName.Like("%" + *filter.Search + "%"),
		)
	}

	if filter.DepartmentID != nil {
		chain = chain.Where(
			generated.Employee.DepartmentID.Eq(*filter.DepartmentID),
		)
	}

	if filter.PositionID != nil {
		chain = chain.Where(
			generated.Employee.PositionID.Eq(*filter.PositionID),
		)
	}

	if filter.Status != nil {
		chain = chain.Where(
			generated.Employee.Status.Eq(string(*filter.Status)),
		)
	}

	return chain.Find(ctx)
}

// Find By NIK
func (r *EmployeeRepository) FindByNIK(
	ctx context.Context,
	nik string,
	preloads ...model.EmployeePreload,
) (*model.Employee, error) {
	employee, err := r.queryWithPreloads(preloads...).
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
	return r.queryWithPreloads(preloads...).
		Where(generated.Employee.DepartmentID.Eq(departmentID)).
		Find(ctx)
}

// Find By Position
func (r *EmployeeRepository) FindAllByPositionID(
	ctx context.Context,
	positionID uint,
	preloads ...model.EmployeePreload,
) ([]model.Employee, error) {

	return r.queryWithPreloads(preloads...).
		Where(generated.Employee.PositionID.Eq(positionID)).
		Find(ctx)
}

// Count
func (r *EmployeeRepository) Count(
	ctx context.Context,
) (int64, error) {

	return r.query.
		Count(ctx, "*")
}

// TODO nanti di hapus
func (r *EmployeeRepository) ProcessInBatches(
	ctx context.Context,
	batchSize int,
	handler func(data []model.Employee, batch int) error,
) error {
	db := r.db.Session(&gorm.Session{
		Context: ctx,
	})
	query := gorm.G[model.Employee](db)
	return query.
		Preload(
			string(model.EmployeePreloadPosition),
			func(db gorm.PreloadBuilder) error {
				db.Select(generated.Position.BaseSalary.Column().Name)
				return nil
			}).
		Where(generated.Employee.Status.Eq(string(model.EmployeeStatusActive))).
		FindInBatches(ctx, batchSize, handler)
}

func (r *EmployeeRepository) ProcessWithoutPayrollInBatches(
	ctx context.Context,
	batchSize int,
	period time.Time,
	handler func(data []model.Employee, batch int) error,
) error {
	db := r.db.Session(&gorm.Session{
		Context: ctx,
	})
	salarySubQuery := gorm.G[model.Salary](db).
		// Where(generated.Salary.EmployeeID.Expr("= employees.id")).
		Where(generated.Salary.EmployeeID.EqExpr(clause.Eq{Value: "employees.id"})).
		Where(generated.Salary.Period.Eq(period)).
		Select("1")

	generated.Salary.EmployeeID.EqExpr(clause.Eq{Value: "employees.id"})

	query := gorm.G[model.Employee](db)
	return query.
		Preload(
			string(model.EmployeePreloadPosition),
			func(db gorm.PreloadBuilder) error {
				db.Select(generated.Position.BaseSalary.Column().Name)
				return nil
			}).
		Where(generated.Employee.Status.Eq(string(model.EmployeeStatusActive))).
		Not("EXISTS (?)", salarySubQuery).
		FindInBatches(ctx, batchSize, handler)
}

// SELECT *
// FROM employees e
// WHERE NOT EXISTS (

//     SELECT 1
//     FROM salaries s

//     WHERE s.employee_id = e.id
//       AND s.period = '2026-07-01'

// )

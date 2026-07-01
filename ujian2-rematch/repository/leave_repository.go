package repository

import (
	"context"
	"time"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type LeaveRepository struct {
	db *gorm.DB
}

func (r *RepositoryManager) Leave() *LeaveRepository {
	return &LeaveRepository{
		db: r.db,
	}
}

func (r *LeaveRepository) preload(
	db *gorm.DB,
	preloads ...model.LeavePreload,
) *gorm.DB {

	for _, preload := range preloads {
		db = db.Preload(string(preload))
	}

	return db
}

func (r *LeaveRepository) Create(
	ctx context.Context,
	leave *model.Leave,
) error {

	return gorm.G[model.Leave](r.db).
		Create(ctx, leave)
}

func (r *LeaveRepository) Update(
	ctx context.Context,
	leave *model.Leave,
) error {

	_, err := gorm.G[model.Leave](r.db).
		Where(generated.Leave.ID.Eq(leave.ID)).
		Updates(ctx, *leave)

	return err
}

func (r *LeaveRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	_, err := gorm.G[model.Leave](r.db).
		Where(generated.Leave.ID.Eq(id)).
		Delete(ctx)

	return err
}

func (r *LeaveRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.LeavePreload,
) (*model.Leave, error) {

	leave, err := gorm.G[model.Leave](
		r.preload(r.db, preloads...),
	).
		Where(generated.Leave.ID.Eq(id)).
		First(ctx)

	return &leave, err
}

func (r *LeaveRepository) FindAll(
	ctx context.Context,
	preloads ...model.LeavePreload,
) ([]model.Leave, error) {

	return gorm.G[model.Leave](
		r.preload(r.db, preloads...),
	).
		Find(ctx)
}

func (r *LeaveRepository) FindAllByEmployeeID(
	ctx context.Context,
	employeeID uint,
	preloads ...model.LeavePreload,
) ([]model.Leave, error) {

	return gorm.G[model.Leave](
		r.preload(r.db, preloads...),
	).
		Where(generated.Leave.EmployeeID.Eq(employeeID)).
		Find(ctx)
}

func (r *LeaveRepository) FindAllByStatus(
	ctx context.Context,
	status model.LeaveStatus,
	preloads ...model.LeavePreload,
) ([]model.Leave, error) {

	return gorm.G[model.Leave](
		r.preload(r.db, preloads...),
	).
		Where(generated.Leave.Status.Eq(string(status))).
		Find(ctx)
}

func (r *LeaveRepository) FindAllByEmployeeAndStatus(
	ctx context.Context,
	employeeID uint,
	status model.LeaveStatus,
	preloads ...model.LeavePreload,
) ([]model.Leave, error) {

	return gorm.G[model.Leave](
		r.preload(r.db, preloads...),
	).
		Where(generated.Leave.EmployeeID.Eq(employeeID)).
		Where(generated.Leave.Status.Eq(string(status))).
		Find(ctx)
}

func (r *LeaveRepository) FindAllByDateRange(
	ctx context.Context,
	startDate time.Time,
	endDate time.Time,
	preloads ...model.LeavePreload,
) ([]model.Leave, error) {

	return gorm.G[model.Leave](
		r.preload(r.db, preloads...),
	).
		Where(generated.Leave.StartDate.Gte(startDate)).
		Where(generated.Leave.EndDate.Lte(endDate)).
		Find(ctx)
}

func (r *LeaveRepository) ExistsOverlappingLeave(
	ctx context.Context,
	employeeID uint,
	startDate time.Time,
	endDate time.Time,
) (bool, error) {

	count, err := gorm.G[model.Leave](r.db).
		Where(generated.Leave.EmployeeID.Eq(employeeID)).
		Where(generated.Leave.StartDate.Lte(endDate)).
		Where(generated.Leave.EndDate.Gte(startDate)).
		Count(ctx, "*")

	return count > 0, err
}

func (r *LeaveRepository) Count(
	ctx context.Context,
) (int64, error) {

	return gorm.G[model.Leave](r.db).
		Count(ctx, "*")
}

func (r *LeaveRepository) CountByStatus(
	ctx context.Context,
	status model.LeaveStatus,
) (int64, error) {

	return gorm.G[model.Leave](r.db).
		Where(generated.Leave.Status.Eq(string(status))).
		Count(ctx, "*")
}

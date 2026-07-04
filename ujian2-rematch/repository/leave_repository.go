package repository

import (
	"context"
	"errors"
	"time"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type LeaveRepository struct {
	db    *gorm.DB
	query gorm.Interface[model.Leave]
}

func (r *RepositoryManager) Leave() *LeaveRepository {
	return &LeaveRepository{
		db:    r.db,
		query: gorm.G[model.Leave](r.db),
	}
}

func (r *LeaveRepository) queryWithPreloads(
	preloads ...model.LeavePreload,
) gorm.ChainInterface[model.Leave] {
	var chain gorm.ChainInterface[model.Leave] = r.query.Scopes()
	for _, preload := range preloads {
		chain = chain.Preload(string(preload), nil)
	}
	return chain
}

func (r *LeaveRepository) Create(
	ctx context.Context,
	leave *model.Leave,
) error {

	return r.query.
		Create(ctx, leave)
}

func (r *LeaveRepository) Update(
	ctx context.Context,
	leave *model.Leave,
) error {

	_, err := r.query.
		Where(generated.Leave.ID.Eq(leave.ID)).
		Updates(ctx, *leave)

	return err
}

func (r *LeaveRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	_, err := r.query.
		Where(generated.Leave.ID.Eq(id)).
		Delete(ctx)

	return err
}

func (r *LeaveRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.LeavePreload,
) (*model.Leave, error) {

	leave, err := r.queryWithPreloads(preloads...).
		Where(generated.Leave.ID.Eq(id)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrLeaveNotFound
	}

	return &leave, err
}

func (r *LeaveRepository) FindAll(
	ctx context.Context,
	preloads ...model.LeavePreload,
) ([]model.Leave, error) {

	return r.queryWithPreloads(preloads...).
		Find(ctx)
}

func (r *LeaveRepository) FindOverlappingLeave(
	ctx context.Context,
	employeeID uint,
	startDate time.Time,
	endDate time.Time,
) (*model.Leave, error) {

	leave, err := r.query.
		Where(generated.Leave.EmployeeID.Eq(employeeID)).
		Where(generated.Leave.StartDate.Lte(endDate)).
		Where(generated.Leave.EndDate.Gte(startDate)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrLeaveNotFound
	}

	return &leave, err
}

func (r *LeaveRepository) CountBetweenDateAndStatus(
	ctx context.Context,
	employeeID uint,
	startDate time.Time,
	endDate time.Time,
	status model.LeaveStatus,
) (int64, error) {
	return r.query.
		Where(generated.Leave.EmployeeID.Eq(employeeID)).
		Where(generated.Leave.StartDate.Lte(endDate)).
		Where(generated.Leave.EndDate.Gte(startDate)).
		Where(generated.Leave.Status.Eq(string(status))).
		Count(ctx, "*")
}

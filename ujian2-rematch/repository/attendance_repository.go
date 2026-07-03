package repository

import (
	"context"
	"time"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db    *gorm.DB
	query gorm.Interface[model.Attendance]
}

func (r *RepositoryManager) Attendance() *AttendanceRepository {
	return &AttendanceRepository{
		db:    r.db,
		query: gorm.G[model.Attendance](r.db),
	}
}

func (r *AttendanceRepository) queryWithPreloads(
	preloads ...model.AttendancePreload,
) gorm.ChainInterface[model.Attendance] {
	var chain gorm.ChainInterface[model.Attendance] = r.query.Scopes()
	for _, preload := range preloads {
		chain = chain.Preload(string(preload), nil)
	}
	return chain
}

func (r *AttendanceRepository) Create(
	ctx context.Context,
	attendance *model.Attendance,
) error {
	return gorm.G[model.Attendance](r.db).
		Create(ctx, attendance)
}

func (r *AttendanceRepository) Update(
	ctx context.Context,
	attendance *model.Attendance,
) error {

	_, err := gorm.G[model.Attendance](r.db).
		Where(generated.Attendance.ID.Eq(attendance.ID)).
		Updates(ctx, *attendance)

	return err
}

func (r *AttendanceRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	_, err := gorm.G[model.Attendance](r.db).
		Where(generated.Attendance.ID.Eq(id)).
		Delete(ctx)

	return err
}

func (r *AttendanceRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.AttendancePreload,
) (*model.Attendance, error) {

	attendance, err := r.queryWithPreloads(preloads...).
		Where(generated.Attendance.ID.Eq(id)).
		First(ctx)

	return &attendance, err
}

func (r *AttendanceRepository) FindAll(
	ctx context.Context,
	preloads ...model.AttendancePreload,
) ([]model.Attendance, error) {

	return r.queryWithPreloads(preloads...).
		Find(ctx)
}

func (r *AttendanceRepository) FindAllByEmployeeID(
	ctx context.Context,
	employeeID uint,
	preloads ...model.AttendancePreload,
) ([]model.Attendance, error) {

	return r.queryWithPreloads(preloads...).
		Where(generated.Attendance.EmployeeID.Eq(employeeID)).
		Find(ctx)
}

func (r *AttendanceRepository) FindByEmployeeAndDate(
	ctx context.Context,
	employeeID uint,
	date time.Time,
	preloads ...model.AttendancePreload,
) (*model.Attendance, error) {

	attendance, err := r.queryWithPreloads(preloads...).
		Where(generated.Attendance.EmployeeID.Eq(employeeID)).
		Where(generated.Attendance.Date.Eq(date)).
		First(ctx)

	return &attendance, err
}

func (r *AttendanceRepository) FindAllByDate(
	ctx context.Context,
	date time.Time,
	preloads ...model.AttendancePreload,
) ([]model.Attendance, error) {

	return r.queryWithPreloads(preloads...).
		Where(generated.Attendance.Date.Eq(date)).
		Find(ctx)
}

func (r *AttendanceRepository) FindAllByStatus(
	ctx context.Context,
	status model.AttendanceStatus,
	preloads ...model.AttendancePreload,
) ([]model.Attendance, error) {

	return r.queryWithPreloads(preloads...).
		Where(generated.Attendance.Status.Eq(string(status))).
		Find(ctx)
}

func (r *AttendanceRepository) ExistsByEmployeeAndDate(
	ctx context.Context,
	employeeID uint,
	date time.Time,
) (bool, error) {

	count, err := gorm.G[model.Attendance](r.db).
		Where(generated.Attendance.EmployeeID.Eq(employeeID)).
		Where(generated.Attendance.Date.Eq(date)).
		Count(ctx, "*")

	return count > 0, err
}

func (r *AttendanceRepository) Count(
	ctx context.Context,
) (int64, error) {

	return gorm.G[model.Attendance](r.db).
		Count(ctx, "*")
}

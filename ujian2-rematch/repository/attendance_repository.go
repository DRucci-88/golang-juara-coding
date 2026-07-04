package repository

import (
	"context"
	"errors"
	"time"
	"ujian2_rematch/dto"
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

func (r *AttendanceRepository) CreateAndCheckIn(
	ctx context.Context,
	attendance *model.Attendance,
) error {
	return r.query.
		Create(ctx, attendance)
}

func (r *AttendanceRepository) CheckOut(
	ctx context.Context,
	employeeID uint,
) error {
	attendance, err := r.FindTodayAttendance(ctx, employeeID)
	if err != nil {
		return nil
	}
	_, err = r.query.
		Where(generated.Attendance.ID.Eq(attendance.ID)).
		Updates(ctx, *attendance)
	return err
}

func (r *AttendanceRepository) FindTodayAttendance(
	ctx context.Context,
	employeeID uint,
	preloads ...model.AttendancePreload,
) (*model.Attendance, error) {

	today := time.Now()

	today = time.Date(
		today.Year(),
		today.Month(),
		today.Day(),
		0, 0, 0, 0,
		today.Location(),
	)

	atteddance, err := r.queryWithPreloads(preloads...).
		Where(generated.Attendance.EmployeeID.Eq(employeeID)).
		// Where(generated.Attendance.Date.EqExpr(clause.Expr{SQL: "CURRENT_DATE"})).
		Where(generated.Attendance.Date.Eq(today)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrAttendanceNotFound
	}

	return &atteddance, err
}

func (r *AttendanceRepository) Update(
	ctx context.Context,
	attendance *model.Attendance,
) error {

	_, err := r.query.
		Where(generated.Attendance.ID.Eq(attendance.ID)).
		Updates(ctx, *attendance)

	return err
}

func (r *AttendanceRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	_, err := r.query.
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
func (r *AttendanceRepository) ExistsByEmployeeAndDate(
	ctx context.Context,
	employeeID uint,
	date time.Time,
) (bool, error) {

	count, err := r.query.
		Where(generated.Attendance.EmployeeID.Eq(employeeID)).
		Where(generated.Attendance.Date.Eq(date)).
		Count(ctx, "*")

	return count > 0, err
}

func (r *AttendanceRepository) Count(
	ctx context.Context,
) (int64, error) {

	return r.query.
		Count(ctx, "*")
}

func (r *AttendanceRepository) SummaryForPayroll(
	ctx context.Context,
	startDate time.Time,
	endDate time.Time,
	employeeID uint,
) (*dto.AttendanceSummary, error) {

	summary := &dto.AttendanceSummary{}

	attendances, err := r.query.
		Where(generated.Attendance.EmployeeID.Eq(employeeID)).
		Where(generated.Attendance.Date.Gte(startDate)).
		Where(generated.Attendance.Date.Lt(endDate)).
		Find(ctx)
	if err != nil {
		return nil, err
	}
	for _, attendance := range attendances {
		switch attendance.Status {
		case model.AttendanceStatusPresent:
			summary.Present++
		case model.AttendanceStatusLate:
			summary.Late++
		case model.AttendanceStatusAbsent:
			summary.Absent++
		}
	}

	// TODO nanti di coba
	// err := r.query.
	// 	Where(generated.Attendance.EmployeeID.Eq(employeeID)).
	// 	Where(generated.Attendance.Date.Gte(startDate)).
	// 	Where(generated.Attendance.Date.Lt(endDate)).
	// 	Select(`
	// 	SUM(CASE WHEN status = 'PRESENT' THEN 1 ELSE 0 END) AS present,
	// 	SUM(CASE WHEN status = 'LATE' THEN 1 ELSE 0 END) AS late,
	// 	SUM(CASE WHEN status = 'ABSENT' THEN 1 ELSE 0 END) AS absent
	// `).
	// 	Scan(ctx, &summary)
	// if err != nil {
	// 	return nil, err
	// }

	return summary, err
}

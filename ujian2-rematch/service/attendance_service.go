package service

import (
	"context"
	"database/sql"
	"time"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type AttendanceService struct {
	repo           *repository.RepositoryManager
	attendanceRepo *repository.AttendanceRepository
}

func NewAttendanceService(
	repo *repository.RepositoryManager,
) *AttendanceService {
	return &AttendanceService{
		repo:           repo,
		attendanceRepo: repo.Attendance(),
	}
}

func (s *AttendanceService) CheckIn(
	ctx context.Context,
	employeeID uint,
) (*model.Attendance, error) {

	now := time.Now()

	status := model.AttendanceStatusPresent
	// TODO Logic jam masuk (telat dan tepat waktu)

	attendance := &model.Attendance{
		EmployeeID: employeeID,
		Date:       now,
		CheckIn:    sql.NullTime{Time: now, Valid: true},
		Status:     status,
	}
	err := s.attendanceRepo.CreateAndCheckIn(ctx, attendance)
	if err != nil {
		return nil, err
	}
	attendanceCreated, err := s.attendanceRepo.FindTodayAttendance(ctx, employeeID, model.AttendancePreloadEmployee)

	return attendanceCreated, err
}

func (s *AttendanceService) CheckOut(
	ctx context.Context,
	employeeID uint,
) (*model.Attendance, error) {
	now := time.Now()

	status := model.AttendanceStatusPresent
	// TODO Logic jam keluar (telat dan tepat waktu)

	attendance, err := s.attendanceRepo.FindTodayAttendance(ctx, employeeID, model.AttendancePreloadEmployee)

	if err != nil {
		return nil, err
	}

	attendance.CheckOut = sql.NullTime{Time: now, Valid: true}
	attendance.Status = status

	err = s.attendanceRepo.CheckOut(ctx, employeeID)
	return attendance, err
}

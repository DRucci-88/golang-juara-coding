package service

import (
	"context"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type LeaveService struct {
	repo      *repository.RepositoryManager
	leaveRepo *repository.LeaveRepository
}

func NewLeaveService(
	repo *repository.RepositoryManager,
) *LeaveService {
	return &LeaveService{
		repo:      repo,
		leaveRepo: repo.Leave(),
	}
}

func (s *LeaveService) Create(
	ctx context.Context,
	employeeID uint,
	dto *dto.LeaveCreateRequest,
) (*model.Leave, error) {

	overlappingLeave, err := s.leaveRepo.FindOverlappingLeave(ctx, employeeID, dto.StartDate, dto.EndDate)

	if err == nil {
		return overlappingLeave, model.ErrLeaveOverlapping
	}

	leave := &model.Leave{
		EmployeeID: employeeID,
		StartDate:  dto.StartDate,
		EndDate:    dto.EndDate,
		Reason:     dto.Reason,
		Status:     model.LeaveStatusPending,
	}

	err = s.leaveRepo.Create(ctx, leave)

	if err != nil {
		return nil, err
	}
	return leave, err
}

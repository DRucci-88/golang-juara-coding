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

func (s *LeaveService) Approval(
	ctx context.Context,
	leaveID int,
	dto *dto.LeaveApprovalRequest,
) (*model.Leave, error) {
	leave, err := s.leaveRepo.FindByID(ctx, uint(leaveID))
	if err != nil {
		return nil, err
	}

	switch leave.Status {
	case model.LeaveStatusApproved:
		return nil, model.ErrLeaveAlreadyApproved
	case model.LeaveStatusRejected:
		return nil, model.ErrLeaveAlreadyRejected
	}

	leave.Status = dto.Status
	if err := s.leaveRepo.Update(ctx, leave); err != nil {
		return nil, err
	}

	return leave, err
}

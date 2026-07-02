package service

import (
	"context"
	"ujian2_rematch/dto"
	"ujian2_rematch/helper"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type EmployeeService struct {
	repo           *repository.RepositoryManager
	employeeRepo   *repository.EmployeeRepository
	departmentRepo *repository.DepartmentRepository
	positionRepo   *repository.PositionRepository
	userRepo       *repository.UserRepository
}

func NewEmployeeService(
	repo *repository.RepositoryManager,
) *EmployeeService {
	return &EmployeeService{
		repo:           repo,
		employeeRepo:   repo.Employee(),
		departmentRepo: repo.Department(),
		positionRepo:   repo.Position(),
		userRepo:       repo.User(),
	}
}

func (s *EmployeeService) Create(
	ctx context.Context,
	dto *dto.EmployeeCreateRequest,
) (*model.Employee, error) {

	if _, err := s.positionRepo.FindByID(ctx, dto.PositionID); err != nil {
		return nil, err
	}

	if _, err := s.departmentRepo.FindByID(ctx, dto.DepartmentID); err != nil {
		return nil, err
	}

	if _, err := s.userRepo.FindByEmail(ctx, dto.Email); err == nil {
		return nil, model.ErrUserEmailAlreadyExists
	}

	password, err := helper.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	var employeeID uint
	err = s.repo.Transaction(ctx, func(repo *repository.RepositoryManager) error {
		user := &model.User{
			Role:     model.UserRoleEmployee,
			Email:    dto.Email,
			Password: password,
		}
		if err := repo.User().Create(ctx, user); err != nil {
			return err
		}

		employee := &model.Employee{
			NIK:          dto.NIK,
			FullName:     dto.FullName,
			Status:       model.EmployeeStatusActive,
			DepartmentID: dto.DepartmentID,
			PositionID:   dto.PositionID,
			UserID:       user.ID,
		}

		if err := repo.Employee().Create(ctx, employee); err != nil {
			return err
		}
		employeeID = employee.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return s.employeeRepo.FindByID(ctx, employeeID, model.EmployeePreloadUser)
}

func (s *EmployeeService) FindAll(
	ctx context.Context,
	filter *dto.EmployeeFilterRequest,
) ([]model.Employee, error) {
	employees, err := s.employeeRepo.FindAll(
		ctx,
		filter,
		model.EmployeePreloadPosition,
		model.EmployeePreloadDepartment,
	)
	return employees, err
}

func (s *EmployeeService) FindByID(
	ctx context.Context,
	id int,
) (*model.Employee, error) {
	return s.employeeRepo.FindByID(
		ctx,
		uint(id),
		model.EmployeePreloadDepartment,
		model.EmployeePreloadPosition,
	)
}

func (s *EmployeeService) Update(
	ctx context.Context,
	id int,
	dto *dto.EmployeeUpdateRequest,
) (*model.Employee, error) {

	_, errPosition := s.positionRepo.FindByID(ctx, dto.PositionID)
	if errPosition != nil {
		return nil, errPosition
	}

	_, errDepartment := s.departmentRepo.FindByID(ctx, dto.DepartmentID)
	if errDepartment != nil {
		return nil, errDepartment
	}

	_, err := s.employeeRepo.Update(ctx, uint(id), &model.Employee{
		NIK:          dto.NIK,
		FullName:     dto.FullName,
		Status:       model.EmployeeStatusActive,
		DepartmentID: dto.DepartmentID,
		PositionID:   dto.PositionID,
	})
	if err != nil {
		return nil, err
	}

	employee, err := s.employeeRepo.FindByID(ctx, uint(id))
	return employee, err
}

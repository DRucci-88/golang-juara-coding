package service

import (
	"context"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type EmployeeService struct {
	repo           *repository.RepositoryManager
	employeeRepo   *repository.EmployeeRepository
	departmentRepo *repository.DepartmentRepository
	positionRepo   *repository.PositionRepository
}

func NewEmployeeService(
	repo *repository.RepositoryManager,
) *EmployeeService {
	return &EmployeeService{
		repo:           repo,
		employeeRepo:   repo.Employee(),
		departmentRepo: repo.Department(),
		positionRepo:   repo.Position(),
	}
}

func (s *EmployeeService) Create(
	ctx context.Context,
	dto *dto.EmployeeCreateRequest,
) (*model.Employee, error) {

	_, errPosition := s.positionRepo.FindByID(ctx, dto.PositionID)
	if errPosition != nil {
		return nil, errPosition
	}

	_, errDepartment := s.departmentRepo.FindByID(ctx, dto.DepartmentID)
	if errDepartment != nil {
		return nil, errDepartment
	}

	employee := model.Employee{
		NIK:          dto.NIK,
		FullName:     dto.FullName,
		Email:        dto.Email,
		Status:       model.EmployeeStatusActive,
		DepartmentID: dto.DepartmentID,
		PositionID:   dto.PositionID,
	}

	err := s.employeeRepo.Create(ctx, &employee)

	return &employee, err
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
		Email:        dto.Email,
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

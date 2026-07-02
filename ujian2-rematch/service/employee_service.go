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

	_, errPosition := s.positionRepo.FindByID(ctx, dto.PositionID)
	if errPosition != nil {
		return nil, errPosition
	}

	_, errDepartment := s.departmentRepo.FindByID(ctx, dto.DepartmentID)
	if errDepartment != nil {
		return nil, errDepartment
	}

	password, errPassword := helper.HashPassword(dto.Password)
	if errPassword != nil {
		return nil, errPassword
	}

	user := model.User{
		Role:     model.UserRoleEmployee,
		Email:    dto.Email,
		Password: password,
	}

	errUser := s.userRepo.Create(ctx, &user)
	if errUser != nil {
		return nil, errUser
	}

	employee := model.Employee{
		NIK:          dto.NIK,
		FullName:     dto.FullName,
		Status:       model.EmployeeStatusActive,
		DepartmentID: dto.DepartmentID,
		PositionID:   dto.PositionID,
		UserID:       user.ID,
	}

	errEmp := s.employeeRepo.Create(ctx, &employee)

	return &employee, errEmp
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

package service

import (
	"context"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type DepartmentService struct {
	repo           *repository.RepositoryManager
	departmentRepo *repository.DepartmentRepository
}

func NewDepartmentService(
	repo *repository.RepositoryManager,
) *DepartmentService {
	return &DepartmentService{
		repo:           repo,
		departmentRepo: repo.Department(),
	}
}

func (s *DepartmentService) Create(
	ctx context.Context,
	dto *dto.DepartmentCreateRequest,
) (*model.Department, error) {
	department := model.Department{
		Name: dto.Name,
		Code: dto.Code,
	}
	err := s.departmentRepo.Create(ctx, &department)
	return &department, err
}

func (s *DepartmentService) FindByID(
	ctx context.Context,
	id int,
) (*model.Department, error) {
	return s.departmentRepo.FindByID(ctx, uint(id))
}

func (s *DepartmentService) FindAll(
	ctx context.Context,
) ([]model.Department, error) {
	return s.departmentRepo.FindAll(ctx)
}

func (s *DepartmentService) Update(
	ctx context.Context,
	id int,
	dto *dto.DepartmentUpdateRequest,
) (*model.Department, error) {

	_, err := s.departmentRepo.Update(ctx, uint(id), &model.Department{
		Name: dto.Name,
		Code: dto.Code,
	})
	if err != nil {
		return nil, err
	}

	department, err := s.departmentRepo.FindByID(ctx, uint(id))

	return department, err
}

func (s *DepartmentService) Delete(
	ctx context.Context,
	id int,
) (*model.Department, error) {

	department, err := s.departmentRepo.FindByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	err = s.departmentRepo.Delete(ctx, uint(id))

	return department, err
}

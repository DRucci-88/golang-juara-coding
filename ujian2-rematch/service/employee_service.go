package service

import (
	"context"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type EmployeeService struct {
	repo         *repository.RepositoryManager
	employeeRepo *repository.EmployeeRepository
}

func NewEmployeeService(
	repo *repository.RepositoryManager,
) *EmployeeService {
	return &EmployeeService{
		repo:         repo,
		employeeRepo: repo.Employee(),
	}
}

func (s *EmployeeService) Create(
	ctx context.Context,

) (*model.Employee, error) {
	return nil, nil
}

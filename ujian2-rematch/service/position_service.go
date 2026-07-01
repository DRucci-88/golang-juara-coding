package service

import (
	"context"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type PositionService struct {
	repo         *repository.RepositoryManager
	positionRepo *repository.PositionRepository
}

func NewPositionService(
	repo *repository.RepositoryManager,
) *PositionService {
	return &PositionService{
		repo:         repo,
		positionRepo: repo.Position(),
	}
}

func (s *PositionService) Create(
	ctx context.Context,
	dto *dto.PositionCreateRequest,
) (*model.Position, error) {
	position := model.Position{
		Title:      dto.Title,
		BaseSalary: dto.BaseSalary,
	}
	err := s.positionRepo.Create(ctx, &position)
	return &position, err
}

func (s *PositionService) FindByID(
	ctx context.Context,
	id int,
) (*model.Position, error) {
	return s.positionRepo.FindByID(ctx, uint(id))
}

func (s *PositionService) FindAll(
	ctx context.Context,
) ([]model.Position, error) {
	return s.positionRepo.FindAll(ctx)
}

func (s *PositionService) Update(
	ctx context.Context,
	id int,
	dto *dto.PositionUpdateRequest,
) (*model.Position, error) {

	_, err := s.positionRepo.Update(ctx, uint(id), &model.Position{
		Title:      dto.Title,
		BaseSalary: dto.BaseSalary,
	})
	if err != nil {
		return nil, err
	}

	position, err := s.positionRepo.FindByID(ctx, uint(id))

	return position, err
}

func (s *PositionService) Delete(
	ctx context.Context,
	id int,
) (*model.Position, error) {

	position, err := s.positionRepo.FindByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	err = s.positionRepo.Delete(ctx, uint(id))

	return position, err
}

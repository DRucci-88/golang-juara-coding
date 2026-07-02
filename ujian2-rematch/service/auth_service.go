package service

import (
	"context"
	"ujian2_rematch/dto"
	"ujian2_rematch/helper"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type AuthService struct {
	repo     *repository.RepositoryManager
	userRepo *repository.UserRepository
}

func NewAuthService(
	repo *repository.RepositoryManager,
) *AuthService {
	return &AuthService{
		repo:     repo,
		userRepo: repo.User(),
	}
}

func (s *AuthService) Login(
	ctx context.Context,
	dto *dto.LoginRequest,
) (*string, error) {
	user, err := s.userRepo.FindByEmail(ctx, dto.Email, model.UserPreloadEmployee)

	if err != nil {
		return nil, model.ErrAuthUnauthorized
	}

	if !helper.CheckPasswordHash(dto.Password, user.Password) {
		return nil, model.ErrAuthUnauthorized
	}

	token, err := helper.GenerateJWT(user.ID, &user.Employee.ID, user.Email, user.Role)

	if err != nil {
		return nil, err
	}

	return &token, err
}

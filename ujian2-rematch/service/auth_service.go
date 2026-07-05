package service

import (
	"context"
	"log"
	"ujian2_rematch/dto"
	"ujian2_rematch/helper"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"
)

type AuthService struct {
	repo                 *repository.RepositoryManager
	userRepo             *repository.UserRepository
	blackListedTokenRepo *repository.BlackListedTokenRepository
}

func NewAuthService(
	repo *repository.RepositoryManager,
) *AuthService {
	return &AuthService{
		repo:                 repo,
		userRepo:             repo.User(),
		blackListedTokenRepo: repo.BlackListedToken(),
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

	log.Printf("User %+v", user)

	token, err := helper.GenerateJWT(user.ID, &user.Employee.ID, user.Email, user.Role)

	if err != nil {
		return nil, err
	}

	return &token, err
}

func (s *AuthService) Logout(
	ctx context.Context,
	authContext *dto.AuthContext,
) error {
	blacklistedToken := &model.BlackListedToken{
		TokenString: authContext.Token,
		ExpiredAt:   authContext.TokenExpiredAt,
	}
	return s.blackListedTokenRepo.Create(ctx, blacklistedToken)
}

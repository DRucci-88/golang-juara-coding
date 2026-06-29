package usecase

import (
	"context"
	"ujian3/domain"
	"ujian3/helper"
)

type authUseCase struct {
	userRepo             domain.UserRepository
	employeeRepo         domain.EmployeeRepository
	blackListedTokenRepo domain.BlackListedTokenRepository
}

func NewAuthUsecase(
	userRepo domain.UserRepository,
	employeeRepo domain.EmployeeRepository,
) domain.AuthUsecase {
	return &authUseCase{
		userRepo:     userRepo,
		employeeRepo: employeeRepo,
	}
}

func (u *authUseCase) Register(c context.Context, dto *domain.RegisterInputDto) (*domain.User, error) {
	user := &domain.User{
		Email: dto.Email,
		Role:  dto.Role,
	}

	// TODO Database Transaction

	hash, errHash := helper.HashPassword(dto.Password)
	if errHash != nil {
		return nil, domain.ErrAuthPasswordHashFailed
	}

	/// Create User
	user, err := u.userRepo.Create(user, hash)
	if err != nil {
		return nil, domain.ErrAuthRegisterFailed
	}

	if dto.Role == domain.RoleEMPLOYEE {
		employee := &domain.Employee{
			UserID:   user.ID,
			NIK:      dto.NIK,
			FullName: dto.FullName,
		}
		employee, err := u.employeeRepo.Create(employee)
		if err != nil {
			return nil, domain.ErrAuthRegisterFailed
		}
		user.Employee = employee
	}

	return user, nil
}

func (u *authUseCase) Login(c context.Context, dto *domain.LoginInputDto) (string, error) {
	user, errUser := u.userRepo.FindByEmail(dto.Email)
	if errUser != nil {
		return "", domain.ErrAuthLoginFailed
	}

	if !helper.CheckPasswordHash(dto.Password, user.Password) {
		return "", domain.ErrAuthLoginFailed
	}

	token, errToken := helper.GenerateJWT(user.ID, user.Email, user.Role)

	if errToken != nil {
		return "", domain.ErrAuthTokenGenerateFailed
	}
	return token, nil
}

func (s *authUseCase) Logout(authContext *domain.AuthContext) error {

	return s.blackListedTokenRepo.Create(authContext.Token, authContext.TokenExpiresAt)
}

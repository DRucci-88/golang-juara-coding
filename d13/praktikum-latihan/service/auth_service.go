package service

import (
	"errors"
	"praktikum/dto"
	"praktikum/helper"
	"praktikum/model"
	"praktikum/repository"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) (*model.User, error)
	Login(req *dto.LoginRequest) (string, error)
	Me(authContext *dto.AuthContext) (*model.User, error)
}

type authServiceImpl struct {
	db       *gorm.DB
	userRepo repository.UserRepository
}

func NewAuthService(
	db *gorm.DB,
	userRepo repository.UserRepository,
) AuthService {
	return &authServiceImpl{
		db:       db,
		userRepo: userRepo,
	}
}

func (s *authServiceImpl) Register(req *dto.RegisterRequest) (*model.User, error) {
	var newUser model.User
	// Cek duplikasi email
	err := s.db.Transaction(func(tx *gorm.DB) error {
		_, errEmail := s.userRepo.FindByEmail(tx, req.Email)
		if errEmail == nil {
			return errors.New("Email sudah terdaftar di sistem")
		}

		// Hash Password
		hash, errHash := helper.HashPassword(req.Password)
		if errHash != nil {
			return errors.New("Gagal memproses kata sandi")
		}

		if req.Role == "" {
			req.Role = "USER" // default role
		}

		newUser = model.User{
			Email:    req.Email,
			Password: hash,
			Role:     req.Role,
		}

		if err := s.userRepo.Create(tx, &newUser); err != nil {
			return err
		}

		return nil
	})
	return &newUser, err
}

func (s *authServiceImpl) Login(req *dto.LoginRequest) (string, error) {

	user, errUser := s.userRepo.FindByEmail(s.db, req.Email)

	if errUser != nil {
		return "", errors.New("Email atau kata sandi salah")
	}

	if !helper.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("Email atau kata sandi salah")
	}

	token, errToken := helper.GenerateJWT(user.ID, user.Email, user.Role)
	if errToken != nil {
		return "", errors.New("Gagal generate token")
	}

	return token, nil

}

func (s *authServiceImpl) Me(authContext *dto.AuthContext) (*model.User, error) {
	return s.userRepo.FindById(s.db, authContext.UserID)
}

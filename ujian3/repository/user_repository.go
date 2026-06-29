package repository

import (
	"time"
	"ujian3/domain"

	"gorm.io/gorm"
)

type UserDB struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"type:varchar(100);not null;uniqueIndex"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:varchar(20);default:'EMPLOYEE'"`

	Employee *EmployeeDB `gorm:"foreignKey:UserID"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserDB) TableName() string { return "users" }

func fromDomainUser(user *domain.User, hashPassword string) *UserDB {
	return &UserDB{
		ID:       user.ID,
		Email:    user.Email,
		Password: hashPassword,
		Role:     user.Role,
	}
}

func toDomainUser(userDB *UserDB) *domain.User {
	return &domain.User{
		ID:        userDB.ID,
		Email:     userDB.Email,
		Password:  userDB.Password,
		Role:      userDB.Role,
		CreatedAt: userDB.CreatedAt,
		UpdatedAt: userDB.UpdatedAt,
		DeletedAt: userDB.DeletedAt.Time,
	}
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *domain.User, hashPassword string) (*domain.User, error) {
	userDB := fromDomainUser(user, hashPassword)

	if _, err := r.FindByEmail(user.Email); err != nil {
		return nil, domain.ErrUserEmailDuplicate
	}

	err := r.db.Create(userDB).Error

	if err != nil {
		return nil, err
	}

	return toDomainUser(userDB), nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var userDB UserDB
	err := r.db.Where("email = ?", email).Find(&userDB).Error
	return toDomainUser(&userDB), err
}

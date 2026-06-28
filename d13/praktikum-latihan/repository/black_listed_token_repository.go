package repository

import (
	"praktikum/model"
	"time"

	"gorm.io/gorm"
)

type BlackListedTokenRepository interface {
	Create(blackListToken *model.BlackListedToken) error
	IsTokenBlacklisted(token string) (bool, error)
}

type blackListedTokenRepositoryImpl struct {
	db *gorm.DB
}

func NewBlackListedTokenRepository(
	db *gorm.DB,
) BlackListedTokenRepository {
	return &blackListedTokenRepositoryImpl{
		db: db,
	}
}

func (r *blackListedTokenRepositoryImpl) Create(blackListToken *model.BlackListedToken) error {
	return r.db.Create(blackListToken).Error
}

func (r *blackListedTokenRepositoryImpl) IsTokenBlacklisted(token string) (bool, error) {
	var count int64
	err := r.db.Model(&model.BlackListedToken{}).
		Where("token_string = ? AND expired_at > ?", token, time.Now()).
		Error

	return count > 0, err
}

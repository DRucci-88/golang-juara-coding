package repository

import (
	"time"
	"ujian3/domain"

	"gorm.io/gorm"
)

type BlackListedTokenDB struct {
	ID          uint      `gorm:"primaryKey"`
	TokenString string    `gorm:"type:text;uniqueIndex;not null"`
	ExpireAt    time.Time `gorm:"not null"`
}

func (BlackListedTokenDB) TableName() string { return "black_listed_tokens" }

type blackListedTokenRepositoryImpl struct {
	db *gorm.DB
}

func NewBlackListedTokenRepository(
	db *gorm.DB,
) domain.BlackListedTokenRepository {
	return &blackListedTokenRepositoryImpl{
		db: db,
	}
}

func (r *blackListedTokenRepositoryImpl) Create(token string, expireAt time.Time) error {
	blackListToken := BlackListedTokenDB{
		TokenString: token,
		ExpireAt:    expireAt,
	}
	return r.db.Create(&blackListToken).Error
}

func (r *blackListedTokenRepositoryImpl) IsTokenBlacklisted(token string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.BlackListedToken{}).
		Where("token_string = ? AND expire_at > ?", token, time.Now()).
		Count(&count).
		Error

	return count > 0, err
}

func (r *blackListedTokenRepositoryImpl) DeleteExpired() (int64, error) {
	result := r.db.
		Where("expire_at < NOW()").
		Delete(&domain.BlackListedToken{})
	return result.RowsAffected, result.Error
}

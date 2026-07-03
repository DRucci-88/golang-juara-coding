package repository

import (
	"context"
	"time"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type BlackListedTokenRepository struct {
	db    *gorm.DB
	query gorm.Interface[model.BlackListedToken]
}

func (r *RepositoryManager) BlackListedToken() *BlackListedTokenRepository {
	return &BlackListedTokenRepository{
		db:    r.db,
		query: gorm.G[model.BlackListedToken](r.db),
	}
}

func (r *BlackListedTokenRepository) Create(
	ctx context.Context,
	blackListedToken *model.BlackListedToken,
) error {
	return r.query.
		Create(ctx, blackListedToken)
}

func (r *BlackListedTokenRepository) IsTokenBlackListed(
	ctx context.Context,
	token string,
) (bool, error) {
	count, err := r.query.
		Where(generated.BlackListedToken.TokenString.Eq(token)).
		Count(ctx, "*")

	return count > 0, err
}

func (r *BlackListedTokenRepository) DeleteExpired(
	ctx context.Context,
) (int, error) {
	return r.query.
		Where(generated.BlackListedToken.ExpiredAt.Lte(time.Now())).
		Delete(ctx)
}

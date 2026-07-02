package repository

import (
	"context"
	"errors"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type UserRepository struct {
	db    *gorm.DB
	query gorm.Interface[model.User]
}

// Factory Method
func (r *RepositoryManager) User() *UserRepository {
	return &UserRepository{
		db:    r.db,
		query: gorm.G[model.User](r.db),
	}
}

func (r *UserRepository) queryWithPreloads(
	preloads ...model.UserPreload,
) gorm.ChainInterface[model.User] {
	chain := r.query.Scopes()
	for _, preload := range preloads {
		chain = chain.Preload(string(preload), nil)
	}
	return chain
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *model.User,
) error {
	return r.query.Create(ctx, user)
}

func (r *UserRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.UserPreload,
) (*model.User, error) {
	user, err := r.queryWithPreloads(preloads...).
		Where(generated.User.ID.Eq(id)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrUserNotFound
	}

	return &user, err
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
	preloads ...model.UserPreload,
) (*model.User, error) {
	user, err := r.queryWithPreloads(preloads...).
		Where(generated.User.Email.Eq(email)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrUserNotFound
	}
	return &user, err
}

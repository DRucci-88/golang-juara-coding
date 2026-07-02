package repository

import (
	"context"
	"errors"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type PositionRepository struct {
	db    *gorm.DB
	query gorm.Interface[model.Position]
}

func (r *RepositoryManager) Position() *PositionRepository {
	return &PositionRepository{
		db:    r.db,
		query: gorm.G[model.Position](r.db),
	}
}

func (r *PositionRepository) queryWithPreloads(
	preloads ...model.PositionPreload,
) gorm.ChainInterface[model.Position] {
	var chain gorm.ChainInterface[model.Position] = r.query.Scopes()
	for _, preload := range preloads {
		chain = chain.Preload(string(preload), nil)
	}
	return chain
}

func (r *PositionRepository) Create(
	ctx context.Context,
	position *model.Position,
) error {
	err := gorm.G[model.Position](r.db).
		Create(ctx, position)
	return err
}

func (r *PositionRepository) Update(
	ctx context.Context,
	id uint,
	position *model.Position,
) (int, error) {

	rows, err := gorm.G[model.Position](r.db).
		Where(generated.Position.ID.Eq(id)).
		Updates(ctx, *position)

	if rows == 0 {
		return rows, model.ErrPositionNotFound
	}

	return rows, err
}

func (r *PositionRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	_, err := gorm.G[model.Position](r.db).
		Where(generated.Position.ID.Eq(id)).
		Delete(ctx)

	return err
}

func (r *PositionRepository) FindByID(
	ctx context.Context,
	id uint,
	preloads ...model.PositionPreload,
) (*model.Position, error) {

	position, err := r.queryWithPreloads(preloads...).
		Where(generated.Position.ID.Eq(id)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrPositionNotFound
	}

	return &position, err
}

func (r *PositionRepository) FindAll(
	ctx context.Context,
	preloads ...model.PositionPreload,
) ([]model.Position, error) {

	return r.queryWithPreloads(preloads...).
		Find(ctx)
}

func (r *PositionRepository) FindByName(
	ctx context.Context,
	title string,
	preloads ...model.PositionPreload,
) (*model.Position, error) {

	position, err := r.queryWithPreloads(preloads...).
		Where(generated.Position.Title.Eq(title)).
		First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, model.ErrPositionNotFound
	}

	return &position, err
}

func (r *PositionRepository) ExistsByName(
	ctx context.Context,
	title string,
) (bool, error) {

	count, err := gorm.G[model.Position](r.db).
		Where(generated.Position.Title.Eq(title)).
		Count(ctx, "*")

	return count > 0, err
}

func (r *PositionRepository) Count(
	ctx context.Context,
) (int64, error) {

	return gorm.G[model.Position](r.db).
		Count(ctx, "*")
}

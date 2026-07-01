package repository

import (
	"context"
	"ujian2_rematch/model"
	"ujian2_rematch/model/generated"

	"gorm.io/gorm"
)

type PositionRepository struct {
	db *gorm.DB
}

func (r *RepositoryManager) Position() *PositionRepository {
	return &PositionRepository{
		db: r.db,
	}
}

func (r *PositionRepository) preload(
	db *gorm.DB,
	preloads ...model.PositionPreload,
) *gorm.DB {

	for _, preload := range preloads {
		db = db.Preload(string(preload))
	}

	return db
}

func (r *PositionRepository) Create(
	ctx context.Context,
	position *model.Position,
) error {

	return gorm.G[model.Position](r.db).
		Create(ctx, position)
}

func (r *PositionRepository) Update(
	ctx context.Context,
	position *model.Position,
) error {

	_, err := gorm.G[model.Position](r.db).
		Where(generated.Position.ID.Eq(position.ID)).
		Updates(ctx, *position)

	return err
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

	position, err := gorm.G[model.Position](
		r.preload(r.db, preloads...),
	).
		Where(generated.Position.ID.Eq(id)).
		First(ctx)

	return &position, err
}

func (r *PositionRepository) FindAll(
	ctx context.Context,
	preloads ...model.PositionPreload,
) ([]model.Position, error) {

	return gorm.G[model.Position](
		r.preload(r.db, preloads...),
	).
		Find(ctx)
}

func (r *PositionRepository) FindByName(
	ctx context.Context,
	title string,
	preloads ...model.PositionPreload,
) (*model.Position, error) {

	position, err := gorm.G[model.Position](
		r.preload(r.db, preloads...),
	).
		Where(generated.Position.Title.Eq(title)).
		First(ctx)

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

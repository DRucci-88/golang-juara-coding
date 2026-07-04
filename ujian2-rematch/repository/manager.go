package repository

import (
	"context"

	"gorm.io/gorm"
)

type RepositoryManager struct {
	db *gorm.DB
}

// Constructor (Google Wire)
func NewRepositoryManager(db *gorm.DB) *RepositoryManager {
	return &RepositoryManager{
		db: db,
	}
}

// Returns underlying database instance.
func (r *RepositoryManager) DB() *gorm.DB {
	return r.db
}

// Clone RepositoryManager with another database instance.
// Usually used internally for Transaction().
func (r *RepositoryManager) WithDB(db *gorm.DB) *RepositoryManager {
	return &RepositoryManager{
		db: db,
	}
}

// Execute repositories inside one database transaction.
func (r *RepositoryManager) Transaction(
	ctx context.Context,
	fn func(repo *RepositoryManager) error,
) error {

	return r.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {

			txRepo := r.WithDB(tx)

			return fn(txRepo)
		})
}

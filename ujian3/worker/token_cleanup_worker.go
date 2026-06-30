package worker

import (
	"log"
	"time"
	"ujian3/domain"

	"gorm.io/gorm"
)

type TokenCleanupWorker struct {
	db   *gorm.DB
	repo domain.BlackListedTokenRepository
}

func NewTokenCleanupWorker(
	db *gorm.DB,
	repo domain.BlackListedTokenRepository,
) *TokenCleanupWorker {
	return &TokenCleanupWorker{
		db:   db,
		repo: repo,
	}
}

func (w *TokenCleanupWorker) Start() {
	ticker := time.NewTicker(10 * time.Minute)

	go func() {

		defer ticker.Stop()

		for range ticker.C {
			count, err := w.repo.DeleteExpired()
			if err != nil {
				log.Println("Token Cleanup Failed", err)
				continue
			}
			log.Println(count)
			log.Printf("Expired blacklist Count [%d]\n", count)
		}
	}()
}

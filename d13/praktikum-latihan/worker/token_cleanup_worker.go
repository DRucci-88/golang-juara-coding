package worker

import (
	"log"
	"praktikum/repository"
	"time"

	"gorm.io/gorm"
)

type TokenCleanupWorker struct {
	db   *gorm.DB
	repo repository.BlackListedTokenRepository
}

func NewTokenCleanupWorker(
	db *gorm.DB,
	repo repository.BlackListedTokenRepository,
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

package worker

import (
	"context"
	"log"
	"time"
	"ujian2_rematch/repository"
)

type TokenCleanupWorker struct {
	repo *repository.RepositoryManager
}

func NewTokenCleanupWorker(
	repo *repository.RepositoryManager,
) *TokenCleanupWorker {
	return &TokenCleanupWorker{
		repo: repo,
	}
}

func (w *TokenCleanupWorker) Start() {
	ticker := time.NewTicker(10 * time.Minute)
	log.Println("Token Cleanup Worker Start")

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			count, err := w.repo.BlackListedToken().DeleteExpired(context.Background())
			if err != nil {
				log.Print("Token Cleanup Failed", err)
			}
			log.Printf("Expired Token [%d]\n", count)
		}
	}()
}

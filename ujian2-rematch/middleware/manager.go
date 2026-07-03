package middleware

import (
	"ujian2_rematch/repository"

	"github.com/gin-gonic/gin"
)

type MiddlewareManager struct {
	JWT gin.HandlerFunc
}

func NewMiddlewareManager(
	repo *repository.RepositoryManager,
) *MiddlewareManager {
	blackListedTokenRepo := repo.BlackListedToken()
	return &MiddlewareManager{
		JWT: NewJWTMiddleware(blackListedTokenRepo),
	}
}

package app

import (
	"praktikum/middleware"
	"praktikum/repository"

	"github.com/gin-gonic/gin"
)

type GroupMiddleware struct {
	JWT gin.HandlerFunc
}

func NewGroupMiddleware(
	blackListedTokenRepo repository.BlackListedTokenRepository,
) *GroupMiddleware {
	return &GroupMiddleware{
		JWT: middleware.NewJWTAuthMiddleware(blackListedTokenRepo),
	}
}

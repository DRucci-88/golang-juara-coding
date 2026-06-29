package app

import (
	"ujian3/domain"
	"ujian3/middleware"

	"github.com/gin-gonic/gin"
)

type GroupMiddleware struct {
	JWT gin.HandlerFunc
}

func NewGroupMiddleware(
	blackListedTokenRepo domain.BlackListedTokenRepository,
) *GroupMiddleware {
	return &GroupMiddleware{
		JWT: middleware.NewJWTAuthMiddleware(blackListedTokenRepo),
	}
}

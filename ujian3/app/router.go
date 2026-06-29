package app

import (
	"ujian3/delivery"
	"ujian3/domain"
	"ujian3/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	m *GroupMiddleware,
	authHandler delivery.AuthHandler,
) *gin.Engine {

	hrdMiddleware := middleware.RequiredRoleMiddleware(domain.RoleHRD)
	r := gin.Default()

	api := r.Group("/api/v1")

	authApi := api.Group("/auth")
	authApi.POST("/register", m.JWT, hrdMiddleware, authHandler.Register)
	authApi.POST("/login", authHandler.Login)
	authApi.POST("/logout", m.JWT, authHandler.Logout)

	return r
}

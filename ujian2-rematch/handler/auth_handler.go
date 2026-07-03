package handler

import (
	"net/http"
	"ujian2_rematch/dto"
	"ujian2_rematch/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(
	authService *service.AuthService,
) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(c.Request.Context(), &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	authContextValue, exist := c.Get("auth")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "akses tidak sah",
			"error":   "Unauthorized",
		})
	}

	authContext := authContextValue.(*dto.AuthContext)

	err := h.authService.Logout(c.Request.Context(), authContext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Logout gagal",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged out successfully",
	})
}

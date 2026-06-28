package handler

import (
	"net/http"
	"praktikum/dto"
	"praktikum/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Me(c *gin.Context)
}

type authHandlerImpl struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandlerImpl{
		authService: authService,
	}
}

func (h *authHandlerImpl) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi pengguna berhasil",
		"data":    user,
	})

}

func (h *authHandlerImpl) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi pengguna berhasil",
		"token":   token,
	})
}

func (h *authHandlerImpl) Me(c *gin.Context) {
	authContextValue, exist := c.Get("auth")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "akses tidak sah",
			"error":   "Unauthorized",
		})
	}

	authContext := authContextValue.(*dto.AuthContext)

	user, err := h.authService.Me(authContext)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal mengambil data profile",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Profile",
		"data": dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Role:  user.Role,
		},
	})
}

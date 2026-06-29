package delivery

import (
	"net/http"
	"strings"
	"ujian3/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CompanyEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	return strings.HasSuffix(
		strings.ToLower(email),
		"@company.co.id",
	)
}

func RegisterRequestValidation(sl validator.StructLevel) {
	req := sl.Current().Interface().(RegisterRequest)

	if req.Role == domain.RoleEMPLOYEE {
		if req.NIK == "" {
			sl.ReportError(
				req.NIK,
				"NIK",
				"nik",
				"required_if_employee",
				"",
			)
		}

		if req.FullName == "" {
			sl.ReportError(
				req.FullName,
				"FullName",
				"full_name",
				"required_if_employee",
				"",
			)
		}
	}
}

type RegisterRequest struct {
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required"`
	Role     domain.Role `json:"role" binding:"oneof=HRD EMPLOYEE"`

	NIK      string `json:"nik"`
	FullName string `json:"full_name"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type authHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(
	authUsecase domain.AuthUsecase,
) AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	_, err := h.authUsecase.Register(
		c.Request.Context(),
		&domain.RegisterInputDto{
			Email:    req.Email,
			Password: req.Password,
			Role:     req.Role,
			NIK:      req.NIK,
			FullName: req.FullName,
		},
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed register", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register Success"})
}

func (h *authHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUsecase.Login(c.Request.Context(), &domain.LoginInputDto{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "failed login", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Login Succeed", "token": token})
}

func (h *authHandler) Logout(c *gin.Context) {
	authContextValue, exist := c.Get("auth")

	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "akses tidak sah",
			"error":   "Unauthorized",
		})
	}

	authContext := authContextValue.(*domain.AuthContext)

	err := h.authUsecase.Logout(authContext)
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

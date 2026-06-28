package middleware

import (
	"net/http"
	"praktikum/dto"

	"github.com/gin-gonic/gin"
)

func RequiredRoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authContextValue, exists := c.Get("auth")

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "akses tidak sah",
				"error":   "Unauthorized",
			})
			return
		}

		authContext := authContextValue.(*dto.AuthContext)

		var next = false
		for _, role := range allowedRoles {
			if role == authContext.Role {
				next = true
				break
			}
		}
		if next {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "akses tidak sah",
				"error":   "Unauthorized",
			})
		}

	}
}

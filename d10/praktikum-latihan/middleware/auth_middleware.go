package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-API-KEY")

		if token != "RAHASIA_NEGARA" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Akses tidak sah, X-API-KEY tidak valid",
			})
		}

		c.Next()
	}
}

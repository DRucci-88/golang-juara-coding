package middleware

import (
	"net/http"
	"strings"
	"ujian3/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func NewJWTAuthMiddleware(
	blackListedTokenRepo domain.BlackListedTokenRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Header otorisasi kosong"})
			return
		}

		// Memisahkan teks "Bearer ..."
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Format otorisasi harus Bearer [token]"})
			return
		}

		tokenString := tokenParts[1]

		// 1. Cek apakah token masuk dalam daftar Blacklist (Logout)
		isTokenBlacklisted, err := blackListedTokenRepo.IsTokenBlacklisted(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate token"})
			return
		}
		if isTokenBlacklisted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is blacklisted"})
			return
		}

		claims := &domain.JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			return domain.JWTSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau telah kedaluwarsa"})
			return
		}

		// Simpan claims ke context
		// c.Set("userID", claims.UserID)
		// c.Set("email", claims.Email)
		// c.Set("role", claims.Role)
		// c.Set("tokenString", tokenString)
		// c.Set("tokenExp", claims.ExpiresAt.Time)

		authContext := domain.AuthContext{
			UserID:         claims.UserID,
			Email:          claims.Email,
			Role:           claims.Role,
			Token:          tokenString,
			TokenExpiresAt: claims.ExpiresAt.Time,
		}

		c.Set("auth", &authContext)

		c.Next()
	}
}

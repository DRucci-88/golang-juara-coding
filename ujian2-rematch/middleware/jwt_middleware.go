package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"ujian2_rematch/dto"
	"ujian2_rematch/model"
	"ujian2_rematch/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func NewJWTMiddleware(
	blackListedTokenRepo *repository.BlackListedTokenRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Println("authHeader" + authHeader)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": model.ErrAuthWrongAuthorizationHeader.Error()})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		log.Println(tokenParts)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": model.ErrAuthWrongAuthorizationHeader.Error()})
			return
		}

		tokenString := tokenParts[1]

		isTokenBlackListed, err := blackListedTokenRepo.IsTokenBlackListed(context.Background(), tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": model.ErrTokenValidatedFailed.Error()})
			return
		}
		if isTokenBlackListed {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": model.ErrTokenIsBlackListed.Error()})
			return
		}

		claims := &dto.JWTClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
			return dto.JWTSecretKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": model.ErrTokenIsExpired.Error()})
			return
		}

		authContext := &dto.AuthContext{
			UserID:         claims.UserID,
			EmployeeID:     claims.EmployeeID,
			Email:          claims.Email,
			Role:           claims.Role,
			Token:          tokenString,
			TokenExpiredAt: claims.ExpiresAt.Time,
		}
		log.Println(authContext)

		c.Set("auth", authContext)
		c.Next()
	}
}

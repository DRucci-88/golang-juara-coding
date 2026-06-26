package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// simple rate limiter (limit 5 request per second per user)
type IPRateLimiter struct {
	mu   sync.Mutex
	hits map[string]int
}

var limiter = IPRateLimiter{
	hits: make(map[string]int),
}

func RateLimiter(maxRequest int) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		limiter.mu.Lock()

		currentHits := limiter.hits[clientIP]

		if currentHits >= maxRequest {
			limiter.mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		limiter.hits[clientIP] = currentHits + 1
		limiter.mu.Unlock()

		c.Next()
	}
}

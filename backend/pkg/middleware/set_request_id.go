package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewString()
		c.Header("X-Request-Id", uuid)
		c.Next()
	}
}
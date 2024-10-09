package middleware

import (
	"context"
	"time"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/config"

	"github.com/gin-gonic/gin"
)

func HttpRequestTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {

		// wrap the request context with a timeout
		timeoutSeconds := config.HttpRequestTimeoutSeconds()
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(timeoutSeconds)*time.Second)
		defer cancel()

		// update gin request context
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

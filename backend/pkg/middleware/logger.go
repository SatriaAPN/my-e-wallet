package middleware

import (
	"time"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/utils"

	core "github.com/SatriaAPN/my-e-wallet/backend/pkg/core"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := utils.GetLogger()
		tn := time.Now()

		requestData := core.NewHttpRequestLogging(c.Request.URL.Path, c.Request.Method, c.Writer.Header().Get("X-Request-Id"), "request")
		logger.Infof(requestData)

		c.Next()

		tp := time.Since(tn)
		responseData := core.NewHttpResponseLogging(c.Request.URL.Path, c.Request.Method, c.Writer.Header().Get("X-Request-Id"), "response", c.Writer.Status(), tp)
		logger.Infof(responseData)
	}
}

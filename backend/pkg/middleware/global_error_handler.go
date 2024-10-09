package middleware

import (
	"net/http"
	"time"

	coreerrors "github.com/SatriaAPN/my-e-wallet/backend/pkg/core/errors"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/utils"

	"github.com/SatriaAPN/my-e-wallet/backend/pkg/core"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()

		if err != nil {
			logger := utils.GetLogger()
			switch e := err.Err.(type) {
			case *errors.Error:
				stackTrace := e.ErrorStack()

				logger.Errorf(core.NewErrorLoggerData("error", c.Writer.Header().Get("X-Request-Id"), stackTrace))

				rCode := http.StatusInternalServerError
				rMessage := e.Error()

				checkErrorStruct(e, &rCode, &rMessage)

				c.AbortWithStatusJSON(rCode, gin.H{
					"error": rMessage,
				})
			case *time.ParseError:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": e.Error(),
				})
			default:
				logger.Errorf(core.NewErrorLoggerData("error", c.Writer.Header().Get("X-Request-Id"), "unhandled"))

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Something Wrong has Happened",
				})
			}
		}
	}
}

func checkErrorStruct(e error, rCode *int, s *string) {
	switch {
	case errors.Is(e, coreerrors.ErrEmailAlreadyExist):
		*rCode = http.StatusBadRequest
	case errors.Is(e, coreerrors.ErrEmailIsNotValid):
		*rCode = http.StatusBadRequest
	case errors.Is(e, coreerrors.ErrWrongPassword):
		*rCode = http.StatusBadRequest
	case errors.Is(e, coreerrors.ErrEmailNotFound):
		*rCode = http.StatusBadRequest
	case errors.Is(e, coreerrors.ErrMinimumPasswordLength):
		*rCode = http.StatusBadRequest
	case errors.Is(e, coreerrors.ErrForgetPasswordTokenLength):
		*rCode = http.StatusBadRequest
	case errors.Is(e, coreerrors.ErrResetCodeNotFound):
		*rCode = http.StatusBadRequest
		*rCode = http.StatusBadRequest
	case errors.Is(e, coreerrors.ErrMaximumPasswordLength):
		*rCode = http.StatusBadRequest
	default:
		*s = "Something Wrong has Happened"
	}
}

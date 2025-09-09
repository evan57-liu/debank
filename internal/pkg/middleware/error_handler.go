package middleware

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/coin50etf/coin-market/internal/pkg/app"
	"github.com/coin50etf/coin-market/internal/pkg/config"
	"github.com/coin50etf/coin-market/internal/pkg/constant"
	"github.com/coin50etf/coin-market/internal/pkg/logger"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // 记录响应内容
	return w.ResponseWriter.Write(b) // 继续向客户端写入
}

// ErrorHandlerMiddleware 处理请求过程中的错误
func ErrorHandlerMiddleware() gin.HandlerFunc {
	var responseWriter *bodyLogWriter
	return func(c *gin.Context) {
		if config.Conf.Log.Level == constant.LogLevelDebug {
			responseWriter = &bodyLogWriter{ResponseWriter: c.Writer, body: bytes.NewBufferString("")}
			c.Writer = responseWriter
		}

		c.Next() // 执行后续 handler

		ctx := c.Request.Context()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var validationErrs validator.ValidationErrors
			if errors.As(err, &validationErrs) {
				errorsMap := make(map[string]string)
				for _, fieldError := range validationErrs {
					errorsMap[fieldError.Field()] = fieldError.Tag()
				}
				logger.Error(ctx, "Invalid request parameters", "errorsMap", errorsMap)
				app.ErrorWithData(c, http.StatusBadRequest, "Invalid request parameters", errorsMap)

				c.Abort()
				return
			}

			var respErr *app.RespError
			if errors.As(err, &respErr) {
				logger.Error(ctx, "Request error", "error", respErr)
				app.Error(c, respErr.Code, respErr.Message)
			} else {
				logger.Error(ctx, "Internal server error", "error", err)
				app.Error(c, http.StatusInternalServerError, "Internal server error")
			}

			c.Abort()
			return
		}

		if config.Conf.Log.Level == constant.LogLevelDebug {
			statusCode := c.Writer.Status()
			responseBody := responseWriter.body.String()
			logger.Info(ctx, "Response success", "status", statusCode, "response", responseBody)
		}
	}
}

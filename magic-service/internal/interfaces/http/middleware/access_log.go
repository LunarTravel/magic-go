package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinLogger 返回适配 gin 的访问日志中间件。
func GinLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		fields := []interface{}{
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"latency_ms", latency.Milliseconds(),
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
		}

		if requestID, exists := c.Get(ContextKeyRequestID); exists {
			fields = append(fields, "request_id", requestID)
		}

		if len(c.Errors) > 0 {
			fields = append(fields, "errors", c.Errors.String())
			logger.Errorw("HTTP request completed with errors", fields...)
		} else {
			logger.Infow("HTTP request completed", fields...)
		}
	}
}

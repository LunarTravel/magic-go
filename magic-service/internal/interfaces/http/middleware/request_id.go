package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	// HeaderRequestID 请求 ID 头。
	HeaderRequestID = "X-Request-ID"
	// ContextKeyRequestID 上下文中请求 ID 的 key。
	ContextKeyRequestID = "request_id"
)

// RequestID 为每个请求生成唯一 ID。
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(HeaderRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set(ContextKeyRequestID, requestID)
		c.Header(HeaderRequestID, requestID)
		c.Next()
	}
}

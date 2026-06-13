// Package handler 提供 HTTP 请求处理器。
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"magic-service/internal/application/kernel"
	"magic-service/internal/interfaces/dto"
)

// HealthHandler 健康检查处理器。
type HealthHandler struct {
	healthService *kernel.HealthService
}

// NewHealthHandler 创建健康检查处理器。
func NewHealthHandler(healthService *kernel.HealthService) *HealthHandler {
	return &HealthHandler{healthService: healthService}
}

// Check 处理健康检查请求。
func (h *HealthHandler) Check(c *gin.Context) {
	status := h.healthService.Check()
	dto.Success(c, gin.H{
		"status": "ok",
		"checks": status,
	})
}

// HelloHandler 探活处理器。
type HelloHandler struct{}

// NewHelloHandler 创建探活处理器。
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// SayHello 处理探活请求。
func (h *HelloHandler) SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from magic-go!",
		"version": "0.1.0",
	})
}

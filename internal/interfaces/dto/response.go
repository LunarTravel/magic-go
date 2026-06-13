// Package dto 定义数据传输对象。
package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse 统一 API 响应格式（与 PHP 保持兼容）。
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 返回成功响应。
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 返回错误响应。
func Error(c *gin.Context, httpStatus int, code int, message string) {
	c.JSON(httpStatus, APIResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// BadRequest 返回 400 错误。
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, 40001, message)
}

// Unauthorized 返回 401 错误。
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, 40101, message)
}

// NotFound 返回 404 错误。
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, 40401, message)
}

// InternalError 返回 500 错误。
func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, 50001, message)
}

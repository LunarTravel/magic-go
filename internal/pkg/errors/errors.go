// Package errors 定义统一错误码。
package errors

// 错误码常量（与 PHP 保持兼容）。
const (
	CodeSuccess       = 0
	CodeInvalidParams = 40001
	CodeUnauthorized  = 40101
	CodeForbidden     = 40301
	CodeNotFound      = 40401
	CodeInternalError = 50001
)

// AppError 应用错误。
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewAppError 创建应用错误。
func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (e *AppError) Error() string {
	return e.Message
}

// ErrInvalidParams 参数错误。
func ErrInvalidParams(message string) *AppError {
	return NewAppError(CodeInvalidParams, message)
}

// ErrUnauthorized 未授权。
func ErrUnauthorized(message string) *AppError {
	return NewAppError(CodeUnauthorized, message)
}

// ErrNotFound 未找到。
func ErrNotFound(message string) *AppError {
	return NewAppError(CodeNotFound, message)
}

// ErrInternal 内部错误。
func ErrInternal(message string) *AppError {
	return NewAppError(CodeInternalError, message)
}

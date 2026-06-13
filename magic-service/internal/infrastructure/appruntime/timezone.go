// Package appruntime 提供应用运行时管理，包括优雅关闭和时区设置。
package appruntime

import (
	"fmt"
	"os"
	"time"
)

const (
	// DefaultTimezone 默认时区。
	DefaultTimezone = "Asia/Shanghai"
)

// SetDefaultProcessTimezone 设置进程默认时区。
func SetDefaultProcessTimezone() error {
	loc, err := time.LoadLocation(DefaultTimezone)
	if err != nil {
		return fmt.Errorf("load timezone %s: %w", DefaultTimezone, err)
	}
	time.Local = loc
	return nil
}

// GracefulShutdownManager 管理优雅关闭。
type GracefulShutdownManager struct {
	shutdownHandlers []ShutdownHandler
}

// ShutdownHandler 定义优雅关闭处理接口。
type ShutdownHandler interface {
	Stop() error
}

// NewGracefulShutdownManager 创建优雅关闭管理器。
func NewGracefulShutdownManager() *GracefulShutdownManager {
	return &GracefulShutdownManager{}
}

// RegisterShutdownHandler 注册关闭处理器。
func (m *GracefulShutdownManager) RegisterShutdownHandler(handler ShutdownHandler) {
	m.shutdownHandlers = append(m.shutdownHandlers, handler)
}

// WaitForShutdownSignal 等待中断信号并执行优雅关闭。
func (m *GracefulShutdownManager) WaitForShutdownSignal() {
	sigCh := make(chan os.Signal, 1)
	// 在实际项目中使用 signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	// 这里简化处理，等待信号
	<-sigCh
	for _, handler := range m.shutdownHandlers {
		_ = handler.Stop()
	}
}

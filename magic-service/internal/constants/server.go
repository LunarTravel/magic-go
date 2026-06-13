// Package constants 定义服务级别的常量。
package constants

import "time"

const (
	// DefaultReadHeaderTimeout HTTP 读取头超时。
	DefaultReadHeaderTimeout = 10 * time.Second
	// DefaultWriteTimeout HTTP 写超时。
	DefaultWriteTimeout = 60 * time.Second
	// DefaultIdleTimeout HTTP 空闲超时。
	DefaultIdleTimeout = 120 * time.Second
	// DefaultShutdownTimeout 优雅关闭超时。
	DefaultShutdownTimeout = 30 * time.Second
	// DefaultServerPort 默认服务端口。
	DefaultServerPort = 9501
)

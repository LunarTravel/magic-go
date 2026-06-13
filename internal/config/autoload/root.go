// Package autoload 定义应用配置结构与类型。
package autoload

// Config 是应用的根配置。
type Config struct {
	Server  ServerConfig  `json:"server"`
	MySQL   MySQLConfig   `json:"mysql"`
	Redis   RedisConfig   `json:"redis"`
	Logging LoggingConfig `json:"logging"`
}

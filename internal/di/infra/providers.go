// Package infra 提供基础设施层依赖注入 Provider。
package infra

import (
	"magic-go/internal/config"
	"magic-go/internal/config/autoload"
	"magic-go/internal/infrastructure/logging"
	mysqlp "magic-go/internal/infrastructure/persistence/mysql"
	redisp "magic-go/internal/infrastructure/persistence/redis"

	"go.uber.org/zap"
)

// ProvideConfig 提供应用配置。
func ProvideConfig() *autoload.Config {
	return config.New()
}

// ProvideServerConfig 从根配置中提取 Server 配置。
func ProvideServerConfig(cfg *autoload.Config) *autoload.ServerConfig {
	return &cfg.Server
}

// ProvideLogger 提供日志记录器。
func ProvideLogger(cfg *autoload.Config) *zap.SugaredLogger {
	return logging.NewFromConfig(cfg.Logging)
}

// ProvideMySQLClient 提供 MySQL 客户端。
func ProvideMySQLClient(cfg *autoload.Config, logger *zap.SugaredLogger) (*mysqlp.Client, error) {
	return mysqlp.NewClient(cfg.MySQL, logger.Named("mysql"))
}

// ProvideRedisClient 提供 Redis 客户端。
func ProvideRedisClient(cfg *autoload.Config, logger *zap.SugaredLogger) (*redisp.Client, error) {
	return redisp.NewClient(cfg.Redis, logger.Named("redis"))
}

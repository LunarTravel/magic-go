package main

import (
	"magic-service/internal/config"
	"magic-service/internal/config/autoload"
	"magic-service/internal/application/kernel"
	"magic-service/internal/interfaces/http/handler"
	mysqlp "magic-service/internal/infrastructure/persistence/mysql"
	redisp "magic-service/internal/infrastructure/persistence/redis"

	"go.uber.org/zap"
)

// 以下是依赖组装辅助函数。
// 后续迁移到 Wire 依赖注入框架时可替换为 wire.Build()。

func configNew() *autoload.Config {
	return config.New()
}

func autoloadConfig() *autoload.Config {
	return config.New()
}

func provideMySQL(cfg *autoload.Config, logger *zap.SugaredLogger) (*mysqlp.Client, error) {
	return mysqlp.NewClient(cfg.MySQL, logger.Named("mysql"))
}

func provideRedis(cfg *autoload.Config, logger *zap.SugaredLogger) (*redisp.Client, error) {
	return redisp.NewClient(cfg.Redis, logger.Named("redis"))
}

func provideHealthService(mysqlClient *mysqlp.Client, redisClient *redisp.Client) *kernel.HealthService {
	_ = mysqlClient
	_ = redisClient
	return kernel.NewHealthService()
}

func provideHealthHandler(healthSvc *kernel.HealthService) *handler.HealthHandler {
	return handler.NewHealthHandler(healthSvc)
}

func provideHelloHandler() *handler.HelloHandler {
	return handler.NewHelloHandler()
}

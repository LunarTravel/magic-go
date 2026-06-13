// Package redis 提供 Redis 客户端封装。
package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"magic-service/internal/config/autoload"
)

// Client Redis 客户端封装。
type Client struct {
	client *redis.Client
}

// NewClient 创建 Redis 客户端。
func NewClient(cfg autoload.RedisConfig, logger *zap.SugaredLogger) (*Client, error) {
	if cfg.Host == "" {
		logger.Info("Redis config not set, skipping connection")
		return nil, nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr(),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	logger.Infof("Redis connected to %s", cfg.Addr())
	return &Client{client: client}, nil
}

// Redis 返回底层的 redis.Client。
func (c *Client) Redis() *redis.Client {
	return c.client
}

// Close 关闭连接。
func (c *Client) Close() error {
	if c.client == nil {
		return nil
	}
	return c.client.Close()
}

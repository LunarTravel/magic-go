// Package mysql 提供 MySQL 客户端封装。
package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"magic-service/internal/config/autoload"
	"go.uber.org/zap"
)

// Client MySQL 客户端。
type Client struct {
	db *sql.DB
}

// NewClient 创建 MySQL 客户端。
func NewClient(cfg autoload.MySQLConfig, logger *zap.SugaredLogger) (*Client, error) {
	if cfg.Host == "" {
		logger.Info("MySQL config not set, skipping connection")
		return nil, nil
	}

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}

	// 连接池配置
	maxOpen := cfg.MaxOpenConns
	if maxOpen <= 0 {
		maxOpen = 100
	}
	maxIdle := cfg.MaxIdleConns
	if maxIdle <= 0 {
		maxIdle = 20
	}
	lifetime := time.Duration(cfg.ConnMaxLifetimeSeconds) * time.Second
	if lifetime <= 0 {
		lifetime = time.Hour
	}

	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	db.SetConnMaxLifetime(lifetime)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping mysql: %w", err)
	}

	logger.Infof("MySQL connected to %s:%d/%s", cfg.Host, cfg.Port, cfg.Database)
	return &Client{db: db}, nil
}

// DB 返回底层的 sql.DB。
func (c *Client) DB() *sql.DB {
	return c.db
}

// Close 关闭连接。
func (c *Client) Close() error {
	if c.db == nil {
		return nil
	}
	return c.db.Close()
}

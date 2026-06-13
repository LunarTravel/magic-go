// Package main 是 magic-go 应用程序的入口点。
package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"magic-go/internal/infrastructure/appruntime"
	"magic-go/internal/infrastructure/logging"
	httpserver "magic-go/internal/interfaces/http"
	"magic-go/internal/pkg/logkey"

	"go.uber.org/zap"
)

func main() {
	// 设置进程时区
	if err := appruntime.SetDefaultProcessTimezone(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to set process timezone: %v\n", err)
		os.Exit(1)
	}

	ctx := context.Background()
	processStartedAt := time.Now()

	// 初始化主 logger
	logger := logging.New()
	defer recoverAndExit(ctx, logger, "magic-go main goroutine panic", processStartedAt)
	logProcessStarting(ctx, logger, processStartedAt)

	// 初始化应用
	initStartedAt := time.Now()
	server, cleanup, err := initializeApplication(logger)
	if err != nil {
		logger.Fatalw("Failed to initialize application",
			logkey.Error, err,
			"pid", os.Getpid(),
			"uptime_seconds", time.Since(processStartedAt).Seconds(),
		)
	}
	logger.Infow("Application initialized",
		logkey.DurationMS, time.Since(initStartedAt).Milliseconds(),
		"pid", os.Getpid(),
	)
	if cleanup != nil {
		defer cleanup()
	}

	// 异步启动 HTTP 服务
	go func() {
		defer recoverAndExit(ctx, logger, "HTTP server goroutine panic", processStartedAt)
		if err := server.Start(ctx); err != nil {
			logger.Fatalw("Failed to start HTTP server",
				logkey.Error, err,
				"pid", os.Getpid(),
			)
		}
	}()

	// 等待关闭信号
	sig := httpserver.WaitForShutdownSignal()
	logger.Infow("Received shutdown signal, shutting down...",
		"signal", sig,
		"pid", os.Getpid(),
	)

	// 优雅关闭
	if err := server.Stop(); err != nil {
		logger.Errorw("Failed to stop server gracefully",
			logkey.Error, err,
		)
		os.Exit(1)
	}
}

// initializeApplication 初始化所有依赖并返回 HTTP 服务器。
func initializeApplication(logger *zap.SugaredLogger) (*httpserver.Server, func(), error) {
	// 加载配置
	cfg := configNew()

	// 组装基础设施
	mysqlClient, err := provideMySQL(cfg, logger)
	if err != nil {
		logger.Warnw("MySQL connection skipped or failed", logkey.Error, err)
	}

	redisClient, err := provideRedis(cfg, logger)
	if err != nil {
		logger.Warnw("Redis connection skipped or failed", logkey.Error, err)
	}

	// 组装应用层
	healthSvc := provideHealthService(mysqlClient, redisClient)

	// 组装接口层
	server := httpserver.NewServer(
		&cfg.Server,
		logger,
		provideHealthHandler(healthSvc),
		provideHelloHandler(),
	)

	cleanup := func() {
		if redisClient != nil {
			_ = redisClient.Close()
		}
		if mysqlClient != nil {
			_ = mysqlClient.Close()
		}
	}

	return server, cleanup, nil
}

func logProcessStarting(_ context.Context, logger *zap.SugaredLogger, startedAt time.Time) {
	logger.Infow("magic-go process starting",
		"pid", os.Getpid(),
		"ppid", os.Getppid(),
		"started_at", startedAt.Format(time.RFC3339),
	)
}

func recoverAndExit(_ context.Context, logger *zap.SugaredLogger, message string, processStartedAt time.Time) {
	if recovered := recover(); recovered != nil {
		logger.Errorw(message,
			"panic", fmt.Sprint(recovered),
			"stack", string(debug.Stack()),
			"pid", os.Getpid(),
			"uptime_seconds", time.Since(processStartedAt).Seconds(),
		)
		os.Exit(1)
	}
}

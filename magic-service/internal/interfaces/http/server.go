// Package httpapi 提供 HTTP 服务的初始化与管理。
package httpapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"magic-service/internal/config/autoload"
	"magic-service/internal/constants"
	"magic-service/internal/interfaces/http/handler"
	"magic-service/internal/interfaces/http/middleware"
	"magic-service/internal/interfaces/http/router"
)

// Server 表示 HTTP 服务。
type Server struct {
	engine     *gin.Engine
	httpServer *http.Server
	config     *autoload.ServerConfig
	logger     *zap.SugaredLogger

	healthHandler *handler.HealthHandler
	helloHandler  *handler.HelloHandler
}

// NewServer 创建 Server 实例。
func NewServer(
	cfg *autoload.ServerConfig,
	logger *zap.SugaredLogger,
	healthHandler *handler.HealthHandler,
	helloHandler *handler.HelloHandler,
) *Server {
	gin.SetMode(resolveGinMode(cfg.Mode))
	engine := gin.New()

	return &Server{
		engine:        engine,
		config:        cfg,
		logger:        logger,
		healthHandler: healthHandler,
		helloHandler:  helloHandler,
	}
}

// Start 启动 HTTP 服务。
func (s *Server) Start(ctx context.Context) error {
	s.setupMiddleware()
	s.setupRoutes()

	port := s.config.Port
	if port <= 0 {
		port = constants.DefaultServerPort
	}

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           s.engine,
		ReadHeaderTimeout: constants.DefaultReadHeaderTimeout,
		WriteTimeout:      constants.DefaultWriteTimeout,
		IdleTimeout:       constants.DefaultIdleTimeout,
	}

	s.logger.Infof("HTTP server starting on :%d (mode=%s, env=%s)", port, s.config.Mode, s.config.Env)

	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

// Stop 优雅关闭 HTTP 服务。
func (s *Server) Stop() error {
	if s.httpServer == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultShutdownTimeout)
	defer cancel()

	s.logger.Info("HTTP server shutting down gracefully...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server: %w", err)
	}
	s.logger.Info("HTTP server stopped")
	return nil
}

func (s *Server) setupMiddleware() {
	// Recovery
	s.engine.Use(gin.Recovery())
	// Request ID
	s.engine.Use(middleware.RequestID())
	// Access Log
	s.engine.Use(middleware.GinLogger(s.logger))
	// CORS
	s.engine.Use(middleware.CORS())
}

func (s *Server) setupRoutes() {
	s.engine.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	router.SetupRoutes(router.Dependencies{
		Engine:        s.engine,
		BasePath:      s.config.BasePath,
		PprofEnabled:  s.config.PprofEnabled,
		HealthHandler: s.healthHandler,
		HelloHandler:  s.helloHandler,
	})
}

// WaitForShutdownSignal 等待系统中断信号。
func WaitForShutdownSignal() os.Signal {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	return <-sigCh
}

func resolveGinMode(mode string) string {
	switch strings.ToLower(strings.TrimSpace(mode)) {
	case "release":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}

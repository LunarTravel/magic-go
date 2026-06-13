// Package router 提供应用路由配置。
package router

import (
	"net/http/pprof"
	"strings"

	"github.com/gin-gonic/gin"

	"magic-go/internal/interfaces/http/handler"
)

// Dependencies 保存路由初始化所需的依赖。
type Dependencies struct {
	Engine        *gin.Engine
	BasePath      string
	PprofEnabled  bool
	HealthHandler *handler.HealthHandler
	HelloHandler  *handler.HelloHandler
}

// SetupRoutes 注册应用的全部路由。
func SetupRoutes(deps Dependencies) {
	// 根路由
	deps.Engine.GET("/health", deps.HealthHandler.Check)

	// API 分组
	base := normalizeBasePath(deps.BasePath)
	api := deps.Engine.Group(base)

	// 探活接口
	if deps.HelloHandler != nil {
		api.GET("/hello", deps.HelloHandler.SayHello)
	}

	// 业务模块路由分组（占位，后续迁移时填充）
	_ = api.Group("/chat")
	_ = api.Group("/agent")
	_ = api.Group("/contact")
	_ = api.Group("/flow")
	_ = api.Group("/knowledge")
	_ = api.Group("/file")

	// pprof 性能分析
	if deps.PprofEnabled {
		registerPprofRoutes(deps.Engine)
	}
}

func registerPprofRoutes(engine *gin.Engine) {
	debug := engine.Group("/debug/pprof")
	debug.GET("/", gin.WrapF(pprof.Index))
	debug.GET("/cmdline", gin.WrapF(pprof.Cmdline))
	debug.GET("/profile", gin.WrapF(pprof.Profile))
	debug.POST("/symbol", gin.WrapF(pprof.Symbol))
	debug.GET("/symbol", gin.WrapF(pprof.Symbol))
	debug.GET("/trace", gin.WrapF(pprof.Trace))
}

func normalizeBasePath(base string) string {
	base = strings.TrimSpace(base)
	if base == "" {
		return "/api/v1"
	}
	if !strings.HasPrefix(base, "/") {
		base = "/" + base
	}
	return strings.TrimRight(base, "/")
}

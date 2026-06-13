// Package app 提供应用层依赖注入 Provider。
package app

import (
	chatservice "magic-go/internal/application/chat/service"
	agentservice "magic-go/internal/application/agent/service"
	contactservice "magic-go/internal/application/contact/service"
	"magic-go/internal/application/kernel"
	"magic-go/internal/interfaces/http/handler"
)

// ProvideHealthService 提供健康检查服务。
func ProvideHealthService() *kernel.HealthService {
	return kernel.NewHealthService()
}

// ProvideChatService 提供聊天应用服务。
func ProvideChatService() *chatservice.ChatService {
	return chatservice.NewChatService()
}

// ProvideAgentService 提供 Agent 应用服务。
func ProvideAgentService() *agentservice.AgentService {
	return agentservice.NewAgentService()
}

// ProvideContactService 提供联系人应用服务。
func ProvideContactService() *contactservice.ContactService {
	return contactservice.NewContactService()
}

// ProvideHealthHandler 提供健康检查处理器。
func ProvideHealthHandler(healthService *kernel.HealthService) *handler.HealthHandler {
	return handler.NewHealthHandler(healthService)
}

// ProvideHelloHandler 提供探活处理器。
func ProvideHelloHandler() *handler.HelloHandler {
	return handler.NewHelloHandler()
}

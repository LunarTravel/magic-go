// Package service 提供 Agent 应用服务。
package service

import "context"

// AgentService Agent 应用服务。
// TODO: 迁移 PHP Agent 业务逻辑到此。
type AgentService struct{}

// NewAgentService 创建 Agent 服务。
func NewAgentService() *AgentService {
	return &AgentService{}
}

// GetAgent 获取代理信息（占位）。
func (s *AgentService) GetAgent(ctx context.Context, agentID string) (any, error) {
	// TODO: 实现获取代理逻辑
	return nil, nil
}

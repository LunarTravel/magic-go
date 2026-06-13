// Package service 提供 Chat 应用服务。
package service

import (
	"context"
)

// ChatService 聊天应用服务。
// TODO: 迁移 PHP Chat 业务逻辑到此。
type ChatService struct{}

// NewChatService 创建聊天服务。
func NewChatService() *ChatService {
	return &ChatService{}
}

// CreateChat 创建聊天会话（占位）。
func (s *ChatService) CreateChat(ctx context.Context, name string) (string, error) {
	// TODO: 实现创建聊天逻辑
	return "", nil
}

// SendMessage 发送消息（占位）。
func (s *ChatService) SendMessage(ctx context.Context, chatID string, content string) error {
	// TODO: 实现发送消息逻辑
	return nil
}

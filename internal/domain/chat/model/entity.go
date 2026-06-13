// Package model 定义 Chat 领域的实体和值对象。
package model

import "time"

// Chat 表示一个聊天会话实体。
type Chat struct {
	ID               string    `json:"id"`
	TopicID          string    `json:"topic_id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	ConversationID   string    `json:"conversation_id"`
	OrganizationCode string    `json:"organization_code"`
	UserID           string    `json:"user_id"`
	AgentID          string    `json:"agent_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Message 表示聊天消息实体。
type Message struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	Role      string    `json:"role"` // user, assistant, system
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

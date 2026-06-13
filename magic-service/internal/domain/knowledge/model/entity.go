// Package model 定义 Knowledge 领域的实体和值对象。
package model

import "time"

// KnowledgeBase 表示知识库实体。
type KnowledgeBase struct {
	ID               string    `json:"id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Type             string    `json:"type"`
	OrganizationCode string    `json:"organization_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Document 表示知识库文档实体。
type Document struct {
	ID               string    `json:"id"`
	Code             string    `json:"code"`
	KnowledgeBaseID  string    `json:"knowledge_base_id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	Status           string    `json:"status"`
	OrganizationCode string    `json:"organization_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Package service 提供 Contact 应用服务。
package service

import "context"

// ContactService 联系人应用服务。
// TODO: 迁移 PHP Contact 业务逻辑到此。
type ContactService struct{}

// NewContactService 创建联系人服务。
func NewContactService() *ContactService {
	return &ContactService{}
}

// GetContact 获取联系人（占位）。
func (s *ContactService) GetContact(ctx context.Context, magicID string) (any, error) {
	// TODO: 实现获取联系人逻辑
	return nil, nil
}

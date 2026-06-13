// Package model 定义 Contact 领域的实体和值对象。
package model

import "time"

// Contact 表示联系人实体。
type Contact struct {
	ID               string    `json:"id"`
	MagicID          string    `json:"magic_id"`
	Type             int       `json:"type"` // 0: AI, 1: Human
	Status           int       `json:"status"`
	CountryCode      string    `json:"country_code"`
	Phone            string    `json:"phone"`
	Email            string    `json:"email"`
	RealName         string    `json:"real_name"`
	OrganizationCode string    `json:"organization_code"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

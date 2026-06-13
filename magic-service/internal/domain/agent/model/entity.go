// Package model 定义 Agent 领域的实体和值对象。
package model

import "time"

// Agent 表示智能代理实体。
type Agent struct {
	ID               string    `json:"id"`
	AgentName        string    `json:"agent_name"`
	RobotName        string    `json:"robot_name"`
	AgentAvatar      string    `json:"agent_avatar"`
	FlowCode         string    `json:"flow_code"`
	Code             string    `json:"code"`
	OrganizationCode string    `json:"organization_code"`
	Description      string    `json:"description"`
	Instructs        []string  `json:"instructs"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

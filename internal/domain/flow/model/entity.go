// Package model 定义 Flow 领域的实体和值对象。
package model

import "time"

// Flow 表示工作流实体。
type Flow struct {
	ID               string    `json:"id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	OrganizationCode string    `json:"organization_code"`
	UserID           string    `json:"user_id"`
	Nodes            []Node    `json:"nodes"`
	Edges            []Edge    `json:"edges"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// Node 表示工作流节点。
type Node struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"position"`
	Config map[string]any `json:"config"`
}

// Edge 表示工作流边。
type Edge struct {
	ID       string `json:"id"`
	Source   string `json:"source"`
	Target   string `json:"target"`
	SourcePort string `json:"source_port"`
	TargetPort string `json:"target_port"`
}

// Package model 定义 File 领域的实体和值对象。
package model

import "time"

// File 表示文件实体。
type File struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Path             string    `json:"path"`
	Size             int64     `json:"size"`
	MimeType         string    `json:"mime_type"`
	StorageType      string    `json:"storage_type"`
	OrganizationCode string    `json:"organization_code"`
	UserID           string    `json:"user_id"`
	CreatedAt        time.Time `json:"created_at"`
}

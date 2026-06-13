// Package shared 定义跨领域共享的值对象。
package shared

// Pagination 分页参数。
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// Offset 返回分页偏移量。
func (p Pagination) Offset() int {
	if p.Page <= 0 {
		return 0
	}
	return (p.Page - 1) * p.Limit()
}

// Limit 返回每页条数（默认 20）。
func (p Pagination) Limit() int {
	if p.PageSize <= 0 {
		return 20
	}
	if p.PageSize > 100 {
		return 100
	}
	return p.PageSize
}

// PaginatedResult 分页结果。
type PaginatedResult[T any] struct {
	Items      []T   `json:"items"`
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalPages int   `json:"total_pages"`
}

// OrganizationContext 组织上下文（多租户）。
type OrganizationContext struct {
	OrganizationCode string `json:"organization_code"`
	ApplicationID    string `json:"application_id"`
	UserID           string `json:"user_id"`
}

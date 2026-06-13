// Package kernel 提供核心应用服务。
package kernel

// HealthService 健康检查应用服务。
type HealthService struct{}

// NewHealthService 创建健康检查服务。
func NewHealthService() *HealthService {
	return &HealthService{}
}

// Check 执行健康检查。
func (s *HealthService) Check() map[string]string {
	status := map[string]string{
		"server": "ok",
	}
	return status
}

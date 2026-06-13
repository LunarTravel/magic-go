package autoload

// ServerConfig 保存 HTTP 服务设置。
type ServerConfig struct {
	Port         int    `json:"port"`
	Mode         string `json:"mode"`
	BasePath     string `json:"base_path"`
	Env          string `json:"env"`
	PprofEnabled bool   `json:"pprof_enabled"`
}

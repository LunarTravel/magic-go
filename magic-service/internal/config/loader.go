// Package config 提供配置加载器，从 YAML 文件和环境变量加载应用配置。
package config

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"

	"magic-service/internal/config/autoload"
)

// New 通过合并文件配置与环境变量创建配置。
func New() *autoload.Config {
	loadDotEnvIfPresent()
	path := resolveDefaultConfigPath()
	if rawPath, ok := os.LookupEnv("CONFIG_FILE"); ok {
		if trimmedPath := strings.TrimSpace(rawPath); trimmedPath != "" {
			path = trimmedPath
		}
	}

	var cfg autoload.Config
	filePath := filepath.Clean(path)
	if data, err := os.ReadFile(filePath); err == nil {
		expanded := expandEnvPlaceholders(string(data))
		var raw map[string]any
		if err := yaml.Unmarshal([]byte(expanded), &raw); err == nil {
			dec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				WeaklyTypedInput: true,
				Result:           &cfg,
			})
			_ = dec.Decode(raw)
		}
	}

	// 为启动阶段常用字段提供合理的默认值
	if cfg.Logging.Level == "" {
		cfg.Logging.Level = autoload.LogLevelInfo
	}
	if cfg.Logging.Format == "" {
		cfg.Logging.Format = autoload.LogFormatJSON
	}

	return &cfg
}

func resolveDefaultConfigPath() string {
	candidates := []string{"configs/config.yaml"}
	if cwd, err := os.Getwd(); err == nil {
		candidates = []string{
			filepath.Join(cwd, "configs", "config.yaml"),
			filepath.Join(cwd, "config.yaml"),
		}
	}
	for _, candidate := range candidates {
		filePath := filepath.Clean(candidate)
		if _, err := os.Stat(filePath); err == nil {
			return filePath
		}
	}
	return filepath.Clean(candidates[0])
}

// expandEnvPlaceholders 将 ${VAR:-default} 和 ${VAR} 替换为环境变量值。
func expandEnvPlaceholders(s string) string {
	// 处理 ${VAR:-default}
	reWithDefault := regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)(?::[-=]|-)([^}]*)}`)
	s = reWithDefault.ReplaceAllStringFunc(s, func(m string) string {
		sub := reWithDefault.FindStringSubmatch(m)
		if len(sub) != 3 {
			return m
		}
		key, def := sub[1], sub[2]
		if v, ok := os.LookupEnv(key); ok && v != "" {
			return v
		}
		return def
	})
	// 处理 ${VAR}
	reSimple := regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)}`)
	s = reSimple.ReplaceAllStringFunc(s, func(m string) string {
		sub := reSimple.FindStringSubmatch(m)
		if len(sub) != 2 {
			return m
		}
		return os.Getenv(sub[1])
	})
	return s
}

// loadDotEnvIfPresent 尝试从 .env 文件加载环境变量。
// 查找顺序：当前工作目录 → 上一级目录（仓库根）。
// 注意：godotenv.Load 不会覆盖已存在的环境变量，
// 因此生产环境（K8s/Docker）通过环境注入的配置优先级最高，.env 仅用于本地开发。
func loadDotEnvIfPresent() {
	var candidates []string
	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates,
			filepath.Join(cwd, ".env"),            // 当前目录（magic-service/）
			filepath.Join(cwd, "..", ".env"),      // 上一级（仓库根）
		)
	}
	for _, p := range candidates {
		if _, err := os.Stat(filepath.Clean(p)); err == nil {
			_ = godotenv.Load(filepath.Clean(p))
			break
		}
	}
}

package autoload

// LoggingConfig 提供应用日志设置。
type LoggingConfig struct {
	Level  LogLevel  `json:"level"`
	Format LogFormat `json:"format"`
}

// LogLevel 枚举支持的日志级别。
type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

// LogFormat 枚举支持的日志格式。
type LogFormat string

const (
	LogFormatJSON LogFormat = "json"
	LogFormatText LogFormat = "text"
)

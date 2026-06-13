// Package logging 提供结构化日志封装。
package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"magic-go/internal/config/autoload"
)

// SugaredLogger 是 zap.SugaredLogger 的类型别名，方便全局使用。
type SugaredLogger = zap.SugaredLogger

// New 创建默认的 SugaredLogger。
func New() *SugaredLogger {
	return NewFromConfig(autoload.LoggingConfig{
		Level:  autoload.LogLevelInfo,
		Format: autoload.LogFormatJSON,
	})
}

// NewFromConfig 根据配置创建 SugaredLogger。
func NewFromConfig(cfg autoload.LoggingConfig) *SugaredLogger {
	level := parseZapLevel(cfg.Level)
	encoder := newJSONEncoder()
	if cfg.Format == autoload.LogFormatText {
		encoder = newTextEncoder()
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		level,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger.Sugar()
}

// Named 返回带有名称子空间的 Logger。
func Named(logger *SugaredLogger, name string) *SugaredLogger {
	return logger.Named(name)
}

func parseZapLevel(level autoload.LogLevel) zapcore.Level {
	switch level {
	case autoload.LogLevelDebug:
		return zapcore.DebugLevel
	case autoload.LogLevelInfo:
		return zapcore.InfoLevel
	case autoload.LogLevelWarn:
		return zapcore.WarnLevel
	case autoload.LogLevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func newJSONEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func newTextEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

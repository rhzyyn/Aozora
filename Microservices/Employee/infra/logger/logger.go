package logger

import (
	"os"
	"strings"

	"github.com/rhzyyn/Aozora/Microservices/Employee/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	z *zap.SugaredLogger
}

// New initializes the zap logger with structured JSON and level-based config.
func New(cfg *config.Config) *Logger {
	// Determine log level from config or environment
	logLevel := zapcore.InfoLevel
	levelStr := strings.ToLower(cfg.App.LogLevel)
	switch levelStr {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	}

	// Define encoder config
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.LevelKey = "severity"
	encoderCfg.MessageKey = "message"
	encoderCfg.CallerKey = "caller"
	encoderCfg.FunctionKey = zapcore.OmitKey
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	// Build core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	// Build logger with service metadata
	zapLogger := zap.New(core).With(
		zap.String("service", cfg.App.Name),
		zap.String("environment", cfg.App.Env),
	)

	return &Logger{
		z: zapLogger.Sugar(),
	}
}

// Info logs structured info messages
func (l *Logger) Info(msg string, fields ...interface{}) {
	l.z.Infow(msg, fields...)
}

// Warn logs structured warnings
func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.z.Warnw(msg, fields...)
}

// Error logs structured error messages
func (l *Logger) Error(msg string, fields ...interface{}) {
	l.z.Errorw(msg, fields...)
}

// Fatal logs a fatal error and exits
func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.z.Fatalw(msg, fields...)
}

// Sync ensures any buffered logs are flushed
func (l *Logger) Sync() {
	_ = l.z.Sync()
}

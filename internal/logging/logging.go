// References:
// 	https://github.com/google/exposure-notifications-server/blob/master/internal/logging/logger.go
package logging

import (
	"context"
	"go.uber.org/zap"
	_ "template/internal/env"
)

type loggerKey struct{}

var fallbackLogger *zap.SugaredLogger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.MessageKey = "message"
	config.EncoderConfig.LevelKey = "severity"
	config.EncoderConfig.TimeKey = "timestamp"
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	if logger, err := config.Build(); err != nil {
		fallbackLogger = zap.NewNop().Sugar()
	} else {
		fallbackLogger = logger.Named("default").Sugar()
	}
}

//noinspection GoUnusedExportedFunction
func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerKey{}).(*zap.SugaredLogger); ok {
		return logger
	}
	return fallbackLogger
}

type Config struct {
	LogLevel string `json:"log_level"`
}

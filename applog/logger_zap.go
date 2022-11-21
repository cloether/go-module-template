// Package applog
//
// References:
//
//	https://github.com/google/exposure-notifications-server/blob/master/internal/logging/logger.go
package applog

import (
	"context"
	_ "github.com/cloether/go-module-template/internal/env"
	"go.uber.org/zap"
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

// noinspection GoUnusedExportedFunction
func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// noinspection GoUnusedExportedFunction
func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerKey{}).(*zap.SugaredLogger); ok {
		return logger
	}
	return fallbackLogger
}

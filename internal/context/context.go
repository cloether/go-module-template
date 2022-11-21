package context

import (
	"context"
	"log"
)

// defaultValue can be set via SetDefault.
//
//goland:noinspection GoUnusedGlobalVariable
var defaultValue string

type contextKey struct{}

// Context wraps context.Context and includes an additional Logger() method.
type Context interface {
	context.Context

	Logger() log.Logger
	Parent() context.Context
	SetParent(ctx context.Context) Context
}

// noinspection GoUnusedExportedFunction
func WithEnv(ctx context.Context, env string) context.Context {
	return context.WithValue(ctx, contextKey{}, env)
}

// noinspection GoUnusedExportedFunction
func FromContext(ctx context.Context) string {
	if value, ok := ctx.Value(contextKey{}).(string); ok {
		return value
	}
	return defaultValue
}

// noinspection GoUnusedExportedFunction
func FromContextWithDefault(ctx context.Context, defaultValue string) string {
	if value, ok := ctx.Value(contextKey{}).(string); ok {
		return value
	}
	return defaultValue
}

// SetDefault sets the package-level global default logger that will be used
// for Background and TODO contexts.
//
// noinspection GoUnusedExportedFunction
func SetDefault(value string) {
	defaultValue = value
}

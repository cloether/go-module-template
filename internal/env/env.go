// Package env
package env

import "context"

type contextKey struct{}

var defaultEnv string

func init() { defaultEnv = "default" }

// noinspection GoUnusedExportedFunction
func WithEnv(ctx context.Context, env string) context.Context {
	return context.WithValue(ctx, contextKey{}, env)
}

// noinspection GoUnusedExportedFunction
func FromContext(ctx context.Context) string {
	if env, ok := ctx.Value(contextKey{}).(string); ok {
		return env
	}
	return defaultEnv
}

// noinspection GoUnusedExportedFunction
func FromContextWithDefault(ctx context.Context, defaultValue string) string {
	if env, ok := ctx.Value(contextKey{}).(string); ok {
		return env
	}
	return defaultValue
}

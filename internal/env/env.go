// environment context
package env

import "context"

type envKey struct{}

var defaultEnv string

func init() { defaultEnv = "default" }

//noinspection GoUnusedExportedFunction
func WithEnv(ctx context.Context, env string) context.Context {
	return context.WithValue(ctx, envKey{}, env)
}

//noinspection GoUnusedExportedFunction
func FromContext(ctx context.Context) string {
	if env, ok := ctx.Value(envKey{}).(string); ok {
		return env
	}
	return defaultEnv
}

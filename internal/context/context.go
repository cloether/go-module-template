package context

import (
	"context"
	"log"
)

// defaultValue can be set via SetDefault.
//
//goland:noinspection GoUnusedGlobalVariable
var defaultValue string

// Context wraps context.Context and includes an additional
// Logger() method.
type Context interface {
	context.Context

	Logger() log.Logger
	Parent() context.Context
	SetParent(ctx context.Context) Context
}

// SetDefault sets the package-level global default logger that
// will be used for Background and TODO contexts.
func SetDefault(value string) {
	defaultValue = value
}

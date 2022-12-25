package server

import (
	"context"

	"github.com/cloether/go-module-template/internal/logger"
)

// Execute is the Server Entry Point
func Execute(_ context.Context, version string) {
	l := logger.GetInstance("main")
	l.Printf("Server: version=%s\n", version)
}

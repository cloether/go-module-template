package server

import (
	"context"

	"github.com/cloether/go-module-template/applog"
)

// Run is the Server Entry Point
func Run() {
	log := applog.FromContext(context.Background())
	log.Debug("Server")
}

package main

import (
	"context"
	"template/internal/applog"
)

// Server Entry Point
func main() {
	log := applog.FromContext(context.Background())
	log.Debug("Server")
}

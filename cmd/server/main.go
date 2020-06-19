package main

import (
	"context"
	"template/internal/logging"
)

// Server Entry Point
func main() {
	log := logging.FromContext(context.Background())
	log.Debug("Server")
}

package server

import "go.uber.org/zap"

// Run is the Server Entry Point
func Run() {
	logger := zap.SugaredLogger{}
	logger.Debug("Server")
}

package logger

import (
	"log"
	"os"
	"sync"
)

type DefaultLogger struct {
	*log.Logger
	filename string
}

var (
	defaultLogger *DefaultLogger
	once          sync.Once
)

// GetInstance create a singleton instance of the hydra logger
//
//goland:noinspection GoUnusedExportedFunction
func GetInstance(path string) *DefaultLogger {
	once.Do(func() {
		var err error
		defaultLogger, err = createLogger(path)
		if err != nil {
			log.Fatal(err)
		}
	})
	return defaultLogger
}

// createLogger creates a new logger instance
func createLogger(path string) (*DefaultLogger, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return nil, err
	}
	logger := log.New(file, "Hydra ", log.Lshortfile)
	return &DefaultLogger{filename: path, Logger: logger}, nil
}

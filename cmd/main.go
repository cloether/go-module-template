package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	version "github.com/cloether/go-module-template"
	"github.com/cloether/go-module-template/cmd/server"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// trap ctrl+c and call cancel on the context
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	// run the command
	server.Execute(ctx, version.Version)
}

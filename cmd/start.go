package cmd

import (
	"context"

	"github.com/cloether/go-module-template/pkg/server"
	"github.com/spf13/cobra"
)

const defaultAddr = "127.0.0.1"
const defaultPort = 8080

type options struct {
	addr string
	port uint16
}

var arguments options

// startCmd represents the `receive` command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Long:  `start server`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		s := server.New(ctx, server.WithAddr(arguments.addr))
		s.Run(ctx, arguments.addr)
	},
}

func init() {
	startCmd.Flags().StringVarP(&arguments.addr, "addr", "a", defaultAddr, "ip address to server on")
	startCmd.Flags().Uint16VarP(&arguments.port, "port", "p", defaultPort, "port to listen on")
	rootCmd.AddCommand(startCmd)
}

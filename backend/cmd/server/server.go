package main

import (
	"github.com/spf13/cobra"
	"steamdeckhomebrew.decktweaks/api/server"
)

var (
	ServerPort *uint16
)

func NewServerCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "start",
		Short:         "Start sever.",
		SilenceErrors: true,
		RunE:          startServer,
	}

	ServerPort = cmd.PersistentFlags().Uint16P("port", "p", 3001, "Server port to listen on")

	return cmd
}

func startServer(cmd *cobra.Command, args []string) error {
	return server.Run(&server.ServerOpts{
		Host: "localhost",
		Port: *ServerPort,
	})
}

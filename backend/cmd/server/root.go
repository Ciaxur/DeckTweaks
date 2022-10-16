package main

import "github.com/spf13/cobra"

func Execute() error {
	rootCmd := &cobra.Command{
		Use:           "server",
		Short:         "Backend API server for DeckTweaks.",
		SilenceErrors: true,
	}
	rootCmd.AddCommand(NewServerCommands())
	return rootCmd.Execute()
}

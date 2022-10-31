package server

import (
	"fmt"

	"steamdeckhomebrew.decktweaks/pkg/config"
)

// Initialize packages required within the server.
func Init() error {
	// Requires configuration to be initialized for persistent storage.
	if err := config.Init(); err != nil {
		return fmt.Errorf("failed to initialize config: %v", err)
	}

	return nil
}

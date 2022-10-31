package config

import (
	"fmt"
	"os"
	"path"
	"time"
)

var (
	CONFIG_PATH       = "/tmp"
	SETTINGS_FILEPATH = path.Join(CONFIG_PATH, "settings.json")
	LoadedConfig      *Configuration // Configuration is nil if none exists.
)

func Init() error {
	// Update the config path to the user's home directory.
	deck_user_homedir := "/home/deck"

	CONFIG_PATH = path.Join(deck_user_homedir, ".config", "decktweaks")

	// Create a configuration directory if non exist.
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		fmt.Printf("[%s] config directory '%s' does not exist\n", time.Now(), CONFIG_PATH)
		if err := os.MkdirAll(CONFIG_PATH, 0755); err != nil {
			return fmt.Errorf("failed to create config directory under %s: %v", CONFIG_PATH, err)
		}
		fmt.Printf("[%s] config directory '%s' created successfuly\n", time.Now(), CONFIG_PATH)
	}

	// Update file paths.
	SETTINGS_FILEPATH = path.Join(CONFIG_PATH, "settings.json")

	// Read in configuration from file, if one exists.
	if err := LoadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	return nil
}

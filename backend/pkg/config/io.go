package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Loads configuration from the config path into LoadedConfig.
func LoadConfig() error {
	// Ensure in-memory config is empty.
	if LoadedConfig == nil {
		LoadedConfig = &Configuration{}
	}

	// Attempt to load configuration if one exists.
	if _, err := os.Stat(SETTINGS_FILEPATH); !os.IsNotExist(err) {
		data, err := os.ReadFile(SETTINGS_FILEPATH)
		if err != nil {
			return fmt.Errorf("failed to read config file: %v", err)
		}

		if err := json.Unmarshal(data, LoadedConfig); err != nil {
			return fmt.Errorf("failed to deserialize config file: %v", err)
		}
	}
	return nil
}

// Saves loaded config to the config path.
func SaveConfig() error {
	if LoadedConfig == nil {
		return fmt.Errorf("failed to save config: in-memory config is nil")
	}

	data, err := json.Marshal(LoadedConfig)
	if err != nil {
		return fmt.Errorf("failed to serialize in-memory config: %v", err)
	}

	if err := os.WriteFile(SETTINGS_FILEPATH, data, 0644); err != nil {
		return fmt.Errorf("failed to write config to '%s': %v", SETTINGS_FILEPATH, err)
	}

	return nil
}

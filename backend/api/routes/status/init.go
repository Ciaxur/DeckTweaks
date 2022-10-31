package status

import (
	"fmt"
	"time"

	"steamdeckhomebrew.decktweaks/pkg/config"
	"steamdeckhomebrew.decktweaks/pkg/system/notification"
	"steamdeckhomebrew.decktweaks/pkg/system/settings"
)

func Init() {
	// Create initial settings instance.
	currentSettings = settings.NewSettings()

	// Sync settings from loaded configurations. Assign a new configuration if
	// none exist.
	if config.LoadedConfig.BatteryMonitor == nil {
		fmt.Printf("[%s] Using new settings\n", time.Now())
		config.LoadedConfig.BatteryMonitor = &currentSettings.BatteryMonitorStatus
	} else {
		fmt.Printf("[%s] Using persistent settings\n", time.Now())
		currentSettings.BatteryMonitorStatus = *config.LoadedConfig.BatteryMonitor
	}

	// Start the notification monitor.
	monitor = notification.NewMonitor(currentSettings)

	// Start the monitor if the loaded configuration enables it.
	if *currentSettings.BatteryMonitorStatus.Enabled {
		monitor.Start()
	}
}

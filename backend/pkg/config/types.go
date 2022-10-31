package config

import "steamdeckhomebrew.decktweaks/pkg/system/settings"

type Configuration struct {
	BatteryMonitor *settings.BatteryMonitorStatus `json:"battery_monitor"`
}

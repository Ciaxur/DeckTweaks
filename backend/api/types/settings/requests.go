package settings

import "steamdeckhomebrew.decktweaks/pkg/system/settings"

type SetSettingsRequest struct {
	BatterySettings settings.BatteryMonitorStatus `json:"battery"`
}

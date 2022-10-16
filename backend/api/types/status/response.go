package status

import "steamdeckhomebrew.decktweaks/pkg/system/settings"

type StatusResponse struct {
	BatteryPercentage int64  `json:"battery_precentage"`
	Message           string `json:"message"`
}

type SettingsResponse struct {
	Settings settings.Settings `json:"settings"`
	Message  string            `json:"message"`
}

type StatusErrorResponse struct {
	Message string `json:"message"`
}

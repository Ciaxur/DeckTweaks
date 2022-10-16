package status

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	types_settings "steamdeckhomebrew.decktweaks/api/types/settings"
	"steamdeckhomebrew.decktweaks/api/types/status"
	"steamdeckhomebrew.decktweaks/pkg/system/notification"
	"steamdeckhomebrew.decktweaks/pkg/system/settings"
)

var (
	currentSettings *settings.Settings
	monitor         *notification.NotificationMonitor
)

func handleGetSettings(w http.ResponseWriter, r *http.Request) {
	status_resp := status.SettingsResponse{}
	status_resp.Settings = *currentSettings

	b, err := json.Marshal(status_resp)
	if err != nil {
		fmt.Printf("Failed to marshal status request: %v\n", err)
		http.Error(w, "Failed to serialize response status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func handleSetSettings(w http.ResponseWriter, r *http.Request) {
	// Verify the request is of json body type.
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "expected json request", http.StatusBadRequest)
		return
	}

	// Parse the request json body.
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("failed to read request body: %v\n", err)
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	var settingsRequest types_settings.SetSettingsRequest
	if err := json.Unmarshal(b, &settingsRequest); err != nil {
		fmt.Printf("failed to de-serialize request body: %v\n", err)
		http.Error(w, fmt.Sprintf("failed to de-serialize request body: %v", err), http.StatusBadRequest)
		return
	}

	// Validate & store optional requested values.
	maxChargeLimit := settingsRequest.BatterySettings.MaxChargeLimit
	minChargeLimit := settingsRequest.BatterySettings.MinChargeLimit

	if maxChargeLimit != nil {
		if err := currentSettings.BatteryMonitorStatus.SetMaxChargeLimit(*maxChargeLimit); err != nil {
			http.Error(w, fmt.Sprintf("max charge limit error: %v", err), http.StatusBadRequest)
		}
	}

	if minChargeLimit != nil {
		if err := currentSettings.BatteryMonitorStatus.SetMinChargeLimit(*minChargeLimit); err != nil {
			http.Error(w, fmt.Sprintf("min charge limit error: %v", err), http.StatusBadRequest)
		}
	}

	// Start the monitor if requested to do so.
	if settingsRequest.BatterySettings.Enabled != nil {
		if *currentSettings.BatteryMonitorStatus.Enabled != *settingsRequest.BatterySettings.Enabled {
			if *settingsRequest.BatterySettings.Enabled {
				fmt.Printf("[%v] Status: Requesting to start battery monitor...\n", time.Now())
				monitor.Start()
			} else {
				fmt.Printf("[%v] Status: Requesting to stop battery monitor...\n", time.Now())
				monitor.Stop()
			}

			*currentSettings.BatteryMonitorStatus.Enabled = *settingsRequest.BatterySettings.Enabled
		}
	}

	// Respond with the current status.
	buff, err := json.Marshal(currentSettings)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to serialize current status: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buff)
}

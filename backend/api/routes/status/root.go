package status

import (
	"github.com/gorilla/mux"
	"steamdeckhomebrew.decktweaks/pkg/system/notification"
	"steamdeckhomebrew.decktweaks/pkg/system/settings"
)

func CreateRoute(r *mux.Router) error {
	Init()
	r.HandleFunc("/settings", handleGetSettings).Methods("GET")
	r.HandleFunc("/settings", handleSetSettings).Methods("POST")
	r.HandleFunc("/battery", handleGetBatteryStatus).Methods("GET")
	return nil
}

func Init() {
	// Create initial settings instance.
	currentSettings = settings.NewSettings()

	// Start the notification monitor.
	monitor = notification.NewMonitor(currentSettings)
}

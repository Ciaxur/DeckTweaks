package telemetry

import "github.com/gorilla/mux"

func CreateRoute(r *mux.Router) error {
	r.HandleFunc("/battery", HandleBatteryTelemetryWebSocket).Methods("GET")
	return nil
}

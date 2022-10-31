package status

import (
	"github.com/gorilla/mux"
)

func CreateRoute(r *mux.Router) error {
	Init()
	r.HandleFunc("/settings", handleGetSettings).Methods("GET")
	r.HandleFunc("/settings", handleSetSettings).Methods("POST")
	r.HandleFunc("/battery", handleGetBatteryStatus).Methods("GET")
	return nil
}

package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"steamdeckhomebrew.decktweaks/api/routes/ping"
	"steamdeckhomebrew.decktweaks/api/routes/status"
	"steamdeckhomebrew.decktweaks/api/routes/telemetry"
)

func InitRoutes(r *mux.Router) error {
	if err := ping.CreateRoute(r.PathPrefix("/ping").Subrouter()); err != nil {
		return fmt.Errorf("create ping route failed: %s", err)
	}
	if err := status.CreateRoute(r.PathPrefix("/status").Subrouter()); err != nil {
		return fmt.Errorf("create status route failed: %s", err)
	}

	if err := telemetry.CreateRoute(r.PathPrefix("/telemetry").Subrouter()); err != nil {
		return fmt.Errorf("create telemetry route failed: %s", err)
	}
	return nil
}

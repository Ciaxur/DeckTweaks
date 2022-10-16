package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"steamdeckhomebrew.decktweaks/api/routes/ping"
	"steamdeckhomebrew.decktweaks/api/routes/status"
)

func InitRoutes(r *mux.Router) error {
	if err := ping.CreateRoute(r.PathPrefix("/ping").Subrouter()); err != nil {
		return fmt.Errorf("create ping route failed: %s", err)
	}
	if err := status.CreateRoute(r.PathPrefix("/status").Subrouter()); err != nil {
		return fmt.Errorf("create status route failed: %s", err)
	}
	return nil
}

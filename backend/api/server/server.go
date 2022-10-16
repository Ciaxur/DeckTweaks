package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"steamdeckhomebrew.decktweaks/api/middlewares"
	"steamdeckhomebrew.decktweaks/api/routes"
)

type ServerOpts struct {
	Port uint16
	Host string
}

func Run(opts *ServerOpts) error {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", opts.Host, opts.Port),
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Second,
	}

	router := mux.NewRouter()

	// Add Middleware.
	router.Use(middlewares.Logger)

	// Add Routes.
	routes.InitRoutes(router)

	// Fallback unknown paths.
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Uh oh!", http.StatusBadRequest)
	})

	http.Handle("/", router)

	// Start server.
	log.Printf("Listening on %s:%d.\n", opts.Host, opts.Port)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}

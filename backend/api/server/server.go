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
	// Initialize server requirements.
	if err := Init(); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", opts.Host, opts.Port),
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Second,
	}

	router := mux.NewRouter()

	// Add Middlewares.
	router.Use(middlewares.DefaultHeaders)
	router.Use(middlewares.Logger)

	// Add on-response path Middlewares.
	router.Use(middlewares.SaveConfigOnResponse)

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

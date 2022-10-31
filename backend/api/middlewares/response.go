package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"steamdeckhomebrew.decktweaks/pkg/config"
)

// Save in-memory configuration on response.
func SaveConfigOnResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		// Check if request was a POST-request, which means internal configurations might
		// be modified, so presist the change.
		if r.Method == "POST" {
			fmt.Printf("[%s] Saving in-memory configuration\n", time.Now())
			if err := config.SaveConfig(); err != nil {
				fmt.Printf("[%s] INTERNAL ERROR: Failed to save configuration: %v\n", time.Now(), err)
			}
		}
	})
}

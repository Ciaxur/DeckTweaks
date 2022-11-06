package telemetry

import (
	"fmt"
	"net/http"
	"time"

	"steamdeckhomebrew.decktweaks/pkg/system/battery"
)

func HandleBatteryTelemetryWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade RAW HTTP connection to a websocket.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("[%s] [URI=%s] failed to upgrade to a websocket connection: %v\n", time.Now(), r.RequestURI, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Stream the state of the battery.
	battTelem := make(chan battery.BatteryState)
	quit := make(chan bool)
	go battery.StreamBatteryState(battTelem, quit)

	for i := 0; i < 10; i++ {
		quit <- false
		t := <-battTelem

		// Write data to WebSocket.
		if err := conn.WriteJSON(t); err != nil {
			http.Error(w, fmt.Sprintf("failed to write JSON data to socket: %v", err), http.StatusInternalServerError)
			break
		}

		// Close connection if there was a failure.
		if t.Error != nil {
			fmt.Printf("[%s] Websocket Battery Telemetry connection failed: %v", time.Now(), t.Error)
			quit <- true
			break
		}
	}

	quit <- true
}

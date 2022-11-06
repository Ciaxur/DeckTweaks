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

	for {
		quit <- false
		t := <-battTelem

		// Write data to WebSocket.
		if err := conn.WriteJSON(t); err != nil {
			fmt.Printf("[%s] Websocket failed to write JSON data: %v\n", time.Now(), err)
			break
		}

		// Close connection if there was a failure.
		if t.Error != nil {
			fmt.Printf("[%s] Websocket Battery Telemetry connection failed: %v", time.Now(), t.Error)
			break
		}
	}

	quit <- true
}

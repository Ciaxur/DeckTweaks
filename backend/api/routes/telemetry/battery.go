package telemetry

import (
	"fmt"
	"net/http"
	"time"
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

	// TODO: write data to the ws connection.
}

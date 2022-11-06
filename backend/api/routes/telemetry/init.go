package telemetry

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Routes: Websocket endpoints.
var (
	upgrader = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // for now, allow any origin.
	}
)

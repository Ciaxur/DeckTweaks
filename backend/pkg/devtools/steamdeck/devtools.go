package steamdeck

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	DevToolsEndpoint = "http://localhost:8080"
)

type DevToolsCommandParams struct {
	JsExpression string `json:"expression"`
	UserGesture  bool   `json:"userGesture"`
	AwaitPromise bool   `json:"awaitPromise"`
}

type DevToolsCommand struct {
	Id     uint64                `json:"id"`
	Method string                `json:"method"`
	Params DevToolsCommandParams `json:"params"`
}

type DevToolsEndpointEntry struct {
	Description          string `json:"description"`
	FrontendURL          string `json:"devtoolsFrontendUrl"`
	Id                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	Url                  string `json:"url"`
	WebsocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

type DevToolsEndpointResponse struct {
	Endpoints []DevToolsEndpointEntry
}

func GetSPWebsocketEntry() (*DevToolsEndpointEntry, error) {
	// Request available hosted websockets.
	resp, err := http.Get(fmt.Sprintf("%s/json", DevToolsEndpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to request endpoints for the SteamDeck Devtools: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshall response body to a struct with expected populated fields.
	devToolEndpoints := DevToolsEndpointResponse{}
	if err := json.Unmarshal(body, &devToolEndpoints.Endpoints); err != nil {
		return nil, fmt.Errorf("failed deserialize devtool endpoint body: %v", err)
	}

	// Extract the SP Websocket.
	for _, entry := range devToolEndpoints.Endpoints {
		if entry.Title == "SP" {
			return &entry, nil
		}
	}

	return nil, fmt.Errorf("failed to find SP entry")
}

func InjectJs(js_expression string) error {
	sp_endpoint, err := GetSPWebsocketEntry()
	if err != nil {
		return fmt.Errorf("failed to get SP websocket: %v", err)
	}

	client, _, err := websocket.DefaultDialer.Dial(sp_endpoint.WebsocketDebuggerUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to establish a websocket connection with '%s': %v", sp_endpoint.WebsocketDebuggerUrl, err)
	}
	defer client.Close()

	// Construct and inovke the evaluation of a given JS expression.
	cmd := DevToolsCommand{
		Id:     1,
		Method: "Runtime.evaluate",
		Params: DevToolsCommandParams{
			JsExpression: js_expression,
			UserGesture:  true,
			AwaitPromise: true,
		},
	}
	if err := client.WriteJSON(cmd); err != nil {
		return fmt.Errorf("failed to invoke javascript: %v", err)
	}

	return nil
}

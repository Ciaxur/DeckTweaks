package status

import (
	"encoding/json"
	"fmt"
	"net/http"

	"steamdeckhomebrew.decktweaks/api/types/status"
	"steamdeckhomebrew.decktweaks/pkg/system/battery"
)

func handleGetBatteryStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Construct the current status.
	batPerc, err := battery.GetPercentage()
	if err != nil {
		fmt.Printf("Internal failure: %s", err)
		errResp := status.StatusErrorResponse{
			Message: "Internal failure: Battery Percentage failed to be obtained.",
		}
		errBytes, _ := json.Marshal(errResp)
		http.Error(w, string(errBytes), http.StatusInternalServerError)
		return
	}

	statusResponse := status.StatusResponse{
		BatteryPercentage: batPerc,
		Message:           "Success.",
	}
	statusBytes, _ := json.Marshal(statusResponse)

	w.Write(statusBytes)
}

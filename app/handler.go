package app

import (
	"data-collector/internal/model"
	"encoding/json"
	"net/http"
)

func getBasicStartupTestData(w http.ResponseWriter, r *http.Request) {
	startupData := []model.BESSData{
		{Timestamp: 12345, Voltage: 120.23, Current: 10.2},
		{Timestamp: 12346, Voltage: 120.01, Current: 10.3},
		{Timestamp: 12347, Voltage: 119.09, Current: 10.8},
		{Timestamp: 12348, Voltage: 119.81, Current: 10.5},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startupData)
}

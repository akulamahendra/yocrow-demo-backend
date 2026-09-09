package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status: "ok",
		Time:   time.Now().UTC().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}

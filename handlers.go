package main

import (
	"encoding/json"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	response := HealthCheck{Status: 200}
	json.NewEncoder(w).Encode(response)
}

func Version(w http.ResponseWriter, r *http.Request) {
	response := VersionApp{Version: "0.1.0"}
	json.NewEncoder(w).Encode(response)
}

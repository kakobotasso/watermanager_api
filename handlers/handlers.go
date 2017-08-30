package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kakobotasso/watermanager/models"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	response := models.HealthCheck{Status: 200}
	json.NewEncoder(w).Encode(response)
}

func Version(w http.ResponseWriter, r *http.Request) {
	response := models.VersionApp{Version: "0.1.0"}
	json.NewEncoder(w).Encode(response)
}

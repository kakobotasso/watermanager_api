package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kakobotasso/watermanager_api/models"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := models.HealthCheck{Status: 200}
	json.NewEncoder(w).Encode(response)
}

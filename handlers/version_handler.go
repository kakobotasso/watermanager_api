package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kakobotasso/watermanager_api/models"
)

func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := models.VersionApp{Version: "0.1.0"}
	json.NewEncoder(w).Encode(response)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, _ := strconv.Atoi(userId)

	response := models.User{
		Id:       id,
		Name:     "Gopher",
		Email:    "gopher@golang.com",
		Cpf:      "123.456.789-00",
		Username: "gopher",
		Password: "123456",
		Token:    "skjdfihs@#nsdj&jsdnfspai239uwe",
	}
	json.NewEncoder(w).Encode(response)
}

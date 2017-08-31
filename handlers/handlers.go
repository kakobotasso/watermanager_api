package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/kakobotasso/watermanager/models"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := models.HealthCheck{Status: 200}
	json.NewEncoder(w).Encode(response)
}

func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	response := models.VersionApp{Version: "0.1.0"}
	json.NewEncoder(w).Encode(response)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if user.Username == "gopher" && user.Password == "123456" {
		response := models.User{
			Id:    123,
			Name:  "Gopher",
			Token: "skjdfihs@#nsdj&jsdnfspai239uwe",
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		response := models.Errors{
			models.Error{
				Key:     "login_incorrect",
				Message: "Username or password incorrect",
			},
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, _ := strconv.Atoi(userId)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

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

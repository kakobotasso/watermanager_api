package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kakobotasso/watermanager/models"
)

func (e Env) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	id, _ := strconv.Atoi(userID)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var user models.User
	var response interface{}

	if err := e.DB.First(&user, id).Error; err != nil {
		response = models.Errors{
			models.Error{
				Key:     "not_found",
				Message: "User not found",
			},
		}
	} else {
		response = &user
	}

	json.NewEncoder(w).Encode(response)
}

func (e Env) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	user.Token = randToken()

	e.DB.Create(&user)

	response := models.User{
		Name:  user.Name,
		Token: user.Token,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

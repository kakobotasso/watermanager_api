package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kakobotasso/watermanager/models"
)

func (e Env) SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var loginUser models.User
	var response interface{}

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

	if err := e.DB.Where(models.User{Username: user.Username, Password: user.Password}).First(&loginUser).Error; err != nil {
		response = models.Errors{
			models.Error{
				Key:     "login_incorrect",
				Message: "Username or password incorrect",
			},
		}
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		response = models.User{
			Name:  loginUser.Name,
			Token: loginUser.Token,
		}
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(response)
}

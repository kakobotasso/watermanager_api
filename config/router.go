package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kakobotasso/watermanager_api/models"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = models.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

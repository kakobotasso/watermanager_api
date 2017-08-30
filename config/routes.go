package config

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kakobotasso/watermanager/handlers"
	"github.com/kakobotasso/watermanager/models"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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

var routes = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/healthcheck",
		handlers.Healthcheck,
	},
	Route{
		"Version",
		"GET",
		"/version",
		handlers.Version,
	},
}

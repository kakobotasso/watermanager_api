package config

import (
	"net/http"

	"github.com/kakobotasso/watermanager/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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

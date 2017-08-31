package config

import (
	"fmt"
	"net/http"

	"github.com/kakobotasso/watermanager/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func prefix() string {
	return "/v1"
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
	Route{
		"User",
		"GET",
		fmt.Sprintf("%s/user/{userId}", prefix()),
		handlers.GetUser,
	},
}

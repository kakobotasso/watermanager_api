package config

import (
	"fmt"
	"net/http"

	"github.com/kakobotasso/watermanager_api/handlers"
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
	Route{"HealthCheck", "GET", "/healthcheck", handlers.Healthcheck},
	Route{"Version", "GET", "/version", handlers.Version},
	Route{"User", "GET", fmt.Sprintf("%s/user/{userId}", prefix()), handlers.GetUser},
	Route{"Create User", "POST", fmt.Sprintf("%s/user", prefix()), handlers.CreateUser},
	Route{"SignIn", "POST", fmt.Sprintf("%s/signin", prefix()), handlers.SignIn},
	Route{"Consumption", "GET", fmt.Sprintf("%s/consumption/{consumption_type}", prefix()), handlers.GetConsumption},
}

package main

import (
	"log"
	"net/http"

	"github.com/kakobotasso/watermanager/config"
)

func main() {
	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}

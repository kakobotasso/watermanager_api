package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kakobotasso/watermanager/config"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := config.NewRouter()
	log.Fatal(http.ListenAndServe(port, router))
}

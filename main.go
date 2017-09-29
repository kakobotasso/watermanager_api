package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kakobotasso/watermanager/handlers"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kakobotasso/watermanager/models"
)

func main() {
	db, err := gorm.Open("sqlite3", "watermanager.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Consumption{})

	env := handlers.Env{
		DB: db,
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/healthcheck", handlers.Healthcheck).Methods("GET").Name("HealthCheck")
	router.HandleFunc("/version", handlers.Version).Methods("GET").Name("Version")
	router.HandleFunc(fmt.Sprintf("%s/user/{userId}", prefix()), env.GetUser).Methods("GET").Name("Get User")
	router.HandleFunc(fmt.Sprintf("%s/user", prefix()), env.CreateUser).Methods("POST").Name("Create User")
	router.HandleFunc(fmt.Sprintf("%s/signin", prefix()), env.SignIn).Methods("POST").Name("Sign In")
	router.HandleFunc(fmt.Sprintf("%s/consumption", prefix()), env.CreateConsumption).Methods("POST").Name("Create Consumption")
	router.HandleFunc(fmt.Sprintf("%s/consumption/{serial}/{consumption_type}", prefix()), env.GetConsumption).Methods("GET").Name("Get Consumption")
	router.HandleFunc(fmt.Sprintf("%s/consumption/average/{serial}/{consumption_type}", prefix()), env.GetConsumptionAverage).Methods("GET").Name("Get Consumption Average")

	log.Fatal(http.ListenAndServe(port, router))
}

func prefix() string {
	return "/v1"
}

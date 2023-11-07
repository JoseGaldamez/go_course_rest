package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JoseGaldamez/go_course_rest/database"
	"github.com/JoseGaldamez/go_course_rest/internal/users"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Basic configuration
	serverAddress := "127.0.0.1:8000"
	_ = godotenv.Load()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	// configure and setup database with gorm
	db := database.GetConfigDatabase()

	// Setup services with database
	userSrv := users.NewService(db, logger)

	// create endpoints with routerMux and with a service
	router := mux.NewRouter()
	users.CreateRouter("/users", router, userSrv)

	// start server
	server := &http.Server{
		Handler: router,
		Addr:    serverAddress,
	}

	log.Println("====> Listening on: " + serverAddress)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoseGaldamez/go_course_rest/internal/users"
	"github.com/gorilla/mux"
)

func main() {

	serverAddress := "127.0.0.1:8000"

	router := mux.NewRouter()
	userSrv := users.NewService()
	users.CreateRouter("/users", router, userSrv)

	server := &http.Server{
		//Handler:      http.TimeoutHandler(router, 5*time.Second, "Timeout!"), // return an error in this time
		Handler: router,
		Addr:    serverAddress,
		// WriteTimeout: 5 * time.Second, // return error when functions ends
		// ReadTimeout:  5 * time.Second,
	}

	fmt.Println("Listening on: " + serverAddress)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JoseGaldamez/go_course_rest/internal/users"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Basic configuration
	serverAddress := "127.0.0.1:8000"
	_ = godotenv.Load()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	// configure and setup database with gorm
	dsn := getDSN()
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()
	_ = db.AutoMigrate(&users.User{})

	// Setup services with database
	userSrv := users.NewService(db, logger)

	// create endpoints with routerMux and with a service
	router := mux.NewRouter()
	users.CreateRouter("/users", router, userSrv)

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

func getDSN() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
}

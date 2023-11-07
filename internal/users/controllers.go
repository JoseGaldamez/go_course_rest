package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MakeCreateUsersController(service Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		log.Println("[!] Getting create user")

		var req CreateRequest

		// transform data of the body into a struct
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRequest{Error: "Invalid request format."})
			return
		}

		// Validate if the request has all the data required
		statusCode, err := BodyCreateValidation(req)
		if err != nil {
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(ErrorRequest{Error: err.Error()})
			return
		}

		// sent information to the service to save them
		user, errService := service.Create(req.FirstName, req.LastName, req.Email, req.Phone)
		if errService != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(ErrorRequest{Error: errService.Error()})
			return
		}

		// response request
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(user)
	}
}

func MakeGetUsersController(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get user")
		json.NewEncoder(w).Encode(map[string]bool{"get": true})
	}
}

func MakeGetAllUsersController(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getall user")
		json.NewEncoder(w).Encode(map[string]bool{"getall": true})
	}
}

func MakeDeleteUsersController(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("delete user")
		json.NewEncoder(w).Encode(map[string]bool{"deleted": true})
	}
}

func MakeUpdateUsersController(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("update user")
		json.NewEncoder(w).Encode(map[string]bool{"updated": true})
	}
}

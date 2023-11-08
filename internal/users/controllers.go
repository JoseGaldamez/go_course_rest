package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeCreateUsersController(service Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
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

func MakeGetAllUsersController(service Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		users, err := service.GetAll()
		if err != nil {
			json.NewEncoder(w).Encode(ErrorRequest{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}

func MakeGetUsersController(service Service) Controller {
	return func(w http.ResponseWriter, request *http.Request) {
		fmt.Println()
		id := mux.Vars(request)["id"]

		user, err := service.Get(id)
		if err != nil {
			json.NewEncoder(w).Encode(ErrorRequest{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func MakeDeleteUsersController(service Service) Controller {
	return func(w http.ResponseWriter, request *http.Request) {
		fmt.Println()
		id := mux.Vars(request)["id"]

		err := service.Delete(id)
		if err != nil {
			json.NewEncoder(w).Encode(ErrorRequest{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"deleted": true})
	}
}

func MakeUpdateUsersController(service Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		var req UpdateRequest

		// transform data of the body into a struct
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRequest{Error: "Invalid request format."})
			return
		}

		if req.FirstName != nil && *req.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRequest{Error: "first_name is a required field."})
			return
		}

		if req.LastName != nil && *req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(ErrorRequest{Error: "last_name is a required field."})
			return
		}

		id := mux.Vars(r)["id"]

		user, err := service.Update(id, req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(ErrorRequest{Error: err.Error()})
			return
		}

		// response request
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

package users

import "github.com/gorilla/mux"

var endpoint Endpoints

func MakeEnpoint(s Service) Endpoints {
	return Endpoints{
		Create: MakeCreateUsersController(s),
		Get:    MakeGetUsersController(s),
		GetAll: MakeGetAllUsersController(s),
		Update: MakeUpdateUsersController(s),
		Delete: MakeDeleteUsersController(s),
	}
}

func CreateRouter(path string, router *mux.Router, s Service) {
	endpoint = MakeEnpoint(s)
	router.HandleFunc(path, endpoint.Create).Methods("POST")
	router.HandleFunc(path, endpoint.GetAll).Methods("GET")
	router.HandleFunc(path+"/{id}", endpoint.Get).Methods("GET")
	router.HandleFunc(path+"/{id}", endpoint.Update).Methods("PATCH")
	router.HandleFunc(path+"/{id}", endpoint.Delete).Methods("DELETE")
}

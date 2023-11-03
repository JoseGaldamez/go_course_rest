package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	port := ":3333"

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/courses", getCourses)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("go /users")
	io.WriteString(w, "This is my users endpoint.\n")
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("go /courses")
	io.WriteString(w, "This is my courses endpoint.\n")
}

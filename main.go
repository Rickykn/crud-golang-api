package main

import (
	"fmt"
	"net/http"

	"github.com/Rickykn/crud-golang-api.git/database"
	"github.com/Rickykn/crud-golang-api.git/handlers"
	"github.com/Rickykn/crud-golang-api.git/middleware"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("CRUD PRODUCT API")
	database.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "API WARUNG 1.0")
	}).Methods(http.MethodGet)

	r.Use(middleware.LogMiddleWare)

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowed)

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}

}

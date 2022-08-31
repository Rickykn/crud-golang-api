package main

import (
	"fmt"
	"net/http"

	"github.com/Rickykn/crud-golang-api.git/database"
	"github.com/Rickykn/crud-golang-api.git/handlers"
	"github.com/Rickykn/crud-golang-api.git/middleware"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang sample crud api")
}

func main() {
	fmt.Println("CRUD PRODUCT API")
	database.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods(http.MethodGet)

	r.Use(middleware.LogMiddleWare)

	r.HandleFunc("/products", handlers.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/{id}", handlers.GetProductById).Methods(http.MethodGet)

	r.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowed)

	fmt.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}

}

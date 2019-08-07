package main

import (
	"github.com/gorilla/mux"
	"github.com/james-work-account/go-rest-api/controllers/api.cars"
	"github.com/james-work-account/go-rest-api/controllers/api.cars.id"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/cars", api_cars.Generic).Methods("GET")
	r.HandleFunc("/api/cars", api_cars.Generic).Methods("POST")
	r.HandleFunc("/api/cars/{id}", api_cars_id.Generic).Methods("GET")
	r.HandleFunc("/api/cars/{id}", api_cars_id.Generic).Methods("PUT")
	r.HandleFunc("/api/cars/{id}", api_cars_id.Generic).Methods("DELETE")

	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

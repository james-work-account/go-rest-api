package api_cars

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/james-work-account/simple-api/models"
	. "github.com/james-work-account/simple-api/utils/constants"
	"github.com/james-work-account/simple-api/utils/db"
	"log"
	"net/http"
)

func Generic(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
	w.Header().Set(ContentType, ApplicationJson)
	switch r.Method {
	case "GET":
		getAllCars(w, r)
	case "POST":
		createCar(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func getAllCars(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.Cars)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)
	car.ID = uuid.New()
	db.Cars = append(db.Cars, car)
	json.NewEncoder(w).Encode(car)
}
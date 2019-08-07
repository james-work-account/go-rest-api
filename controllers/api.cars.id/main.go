package api_cars_id

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/james-work-account/simple-api/models"
	. "github.com/james-work-account/simple-api/utils/constants"
	"github.com/james-work-account/simple-api/utils/db"
	"log"
	"net/http"
)

func Generic(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
	w.Header().Set(ContentType, ApplicationJson)
	params := mux.Vars(r)

	id, err := uuid.Parse(params["id"])

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{Code: "BAD_REQUEST", Message: "Invalid UUID for ID parameter"})
	} else {
		for i, car := range db.Cars {
			if car.ID == id {
				switch r.Method {
				case "GET":
					getSingleCar(w, r, car)
				case "PUT":
					updateSingleCar(w, r, i, id)
				case "DELETE":
					deleteSingleCar(w, r, i)
				default:
					w.WriteHeader(http.StatusNotFound)
				}
				return
			}
		}
		// Car not found
		w.WriteHeader(http.StatusNoContent)
	}

}

func getSingleCar(w http.ResponseWriter, r *http.Request, car models.Car) {
	json.NewEncoder(w).Encode(car)
	return
}

func updateSingleCar(w http.ResponseWriter, r *http.Request, i int, id uuid.UUID) {
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)
	car.ID = id
	carRef := &db.Cars[i]
	*carRef = car
	json.NewEncoder(w).Encode(db.Cars)
	return
}

func deleteSingleCar(w http.ResponseWriter, r *http.Request, i int) {
	db.Cars = append(db.Cars[:i], db.Cars[i+1:]...)
	json.NewEncoder(w).Encode(db.Cars)
	return
}

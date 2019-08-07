package db

import (
	"github.com/google/uuid"
	"github.com/james-work-account/simple-api/models"
)

var Cars []models.Car = []models.Car{{ID: uuid.New(), Make: "BMW", Model: "X5", Age: 4}, {ID: uuid.New(), Make: "Mercedes Benz", Model: "A Class", Age: 2}}
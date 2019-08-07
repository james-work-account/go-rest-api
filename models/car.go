package models

import "github.com/google/uuid"

type Car struct {
	ID    uuid.UUID `json:"id"`
	Make  string    `json:"make"`
	Model string    `json:"model"`
	Age   int       `json:"age"`
}
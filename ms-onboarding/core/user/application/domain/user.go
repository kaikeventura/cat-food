package domain

import "github.com/google/uuid"

type User struct {
	Identifier uuid.UUID `json:"identifier"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	BirthDate  string    `json:"birth_date"`
	Document   string    `json:"document"`
	UserType   string    `json:"user_type"`
	Status     string    `json:"status"`
	Addresses  []Address `json:"addresses"`
}

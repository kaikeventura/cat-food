package domain

import "github.com/google/uuid"

type Address struct {
	Identifier uuid.UUID `json:"identifier"`
	Street     string    `json:"street"`
	Number     string    `json:"number"`
	District   string    `json:"district"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
}

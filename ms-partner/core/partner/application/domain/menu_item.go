package domain

import "github.com/google/uuid"

type MenuItem struct {
	Identifier  uuid.UUID `json:"identifier"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	IsActive    bool      `json:"is_active"`
}

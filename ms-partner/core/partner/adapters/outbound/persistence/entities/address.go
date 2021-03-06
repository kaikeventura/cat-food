package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Id         uint      `gorm:"primaryKey; autoIncrement"`
	Identifier uuid.UUID `gorm:"not null; unique"`
	Street     string    `gorm:"not null"`
	Number     string    `gorm:"not null"`
	District   string    `gorm:"not null"`
	City       string    `gorm:"not null"`
	State      string    `gorm:"not null"`
	PostalCode string    `gorm:"not null"`
	PartnerId  uint
}

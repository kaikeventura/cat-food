package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id         uint      `gorm:"primaryKey; autoIncrement"`
	Identifier uuid.UUID `gorm:"not null"`
	FirstName  string    `gorm:"not null"`
	LastName   string    `gorm:"not null"`
	BirthDate  string    `gorm:"not null"`
	Document   string    `gorm:"unique; not null"`
	UserType   string    `gorm:"not null"`
	Addresses  []Address `gorm:"foreignKey:UserId"`
}

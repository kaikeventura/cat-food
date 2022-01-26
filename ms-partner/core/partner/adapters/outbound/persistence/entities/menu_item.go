package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuItem struct {
	gorm.Model
	Id          uint      `gorm:"primaryKey; autoIncrement"`
	Identifier  uuid.UUID `gorm:"not null; unique"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Price       int       `gorm:"not null"`
	IsActive    bool      `gorm:"not null"`
	MenuId      uint
}

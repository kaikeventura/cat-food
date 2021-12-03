package entities

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Id         uint   `gorm:"primaryKey; autoIncrement"`
	Street     string `gorm:"not null"`
	Number     string `gorm:"not null"`
	District   string `gorm:"not null"`
	City       string `gorm:"not null"`
	State      string `gorm:"not null"`
	PostalCode string `gorm:"not null"`
	UserId     uint
}
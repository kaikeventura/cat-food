package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint      `gorm:"primaryKey; autoIncrement"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	BirthDate string    `gorm:"not null"`
	Document  string    `gorm:"unique; not null"`
	UserType  string    `gorm:"not null"`
	Address   []Address `gorm:"foreignKey:UserId"`
}

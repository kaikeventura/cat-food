package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	Id             uint      `gorm:"primaryKey; autoIncrement"`
	Identifier     uuid.UUID `gorm:"not null; unique"`
	Name           string    `gorm:"not null"`
	PartnerType    string    `gorm:"not null"`
	UserIdentifier string    `gorm:"not null"`
	Address        Address   `gorm:"foreignKey:PartnerId"`
}

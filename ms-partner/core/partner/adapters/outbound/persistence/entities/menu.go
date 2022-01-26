package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Id         uint      `gorm:"primaryKey; autoIncrement"`
	Identifier uuid.UUID `gorm:"not null; unique"`
	PartnerId  uint
	MenuItems  []MenuItem `gorm:"foreignKey:MenuId"`
}

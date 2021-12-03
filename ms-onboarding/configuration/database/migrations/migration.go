package migrations

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/outbound/persistence/entities"
	"gorm.io/gorm"
	"log"
)

func ExecuteMigrations(database *gorm.DB) {
	err := database.AutoMigrate(entities.User{}, entities.Address{})

	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}

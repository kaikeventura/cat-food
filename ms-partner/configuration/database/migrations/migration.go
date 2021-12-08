package migrations

import (
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"gorm.io/gorm"
	"log"
)

func ExecuteMigrations(database *gorm.DB) {
	err := database.AutoMigrate(entities.Partner{}, entities.Address{})

	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}

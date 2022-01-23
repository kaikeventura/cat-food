package migrations

import (
	"log"

	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"gorm.io/gorm"
)

func ExecuteMigrations(database *gorm.DB) {
	err := database.AutoMigrate(entities.Partner{}, entities.Address{}, entities.Menu{})

	if err != nil {
		log.Fatal("Migration error: ", err)
	}
}

package persistence

import (
	"log"

	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/utils/converters"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
	"gorm.io/gorm"
)

var database *gorm.DB

type MySQLMenuPersistence struct {
}

func ConstructMySQLMenuPersistence(databaseRepository *gorm.DB) MySQLMenuPersistence {
	database = databaseRepository

	return MySQLMenuPersistence{}
}

func (persistence MySQLMenuPersistence) SaveMenu(menu domain.Menu) (domain.Menu, error) {
	menuEntity, err := buildMenuEntity(menu)

	if err != nil {
		return domain.Menu{}, err
	}

	err = database.Create(&menuEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return domain.Menu{}, err
	}

	return converters.MenuEntityToMenuDomain(menuEntity, menu.PartnerIdentifier), nil
}

func buildMenuEntity(menu domain.Menu) (entities.Menu, error) {
	identifier, _ := uuid.NewRandom()
	partner, err := findPartnerByIdentifier(menu.PartnerIdentifier)

	if err != nil {
		return entities.Menu{}, err
	}

	return entities.Menu{
		Identifier: identifier,
		PartnerId:  partner.Id,
	}, err
}

func findPartnerByIdentifier(identifier string) (entities.Partner, error) {
	var partner entities.Partner
	err := database.First(&partner, "partners.identifier = ?", identifier).Error

	if err != nil {
		log.Print("Find partner error: " + err.Error())

		return entities.Partner{}, err
	}

	return partner, nil
}

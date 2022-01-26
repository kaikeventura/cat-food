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

func (persistence MySQLMenuPersistence) SaveMenuItem(menuIdentifier uuid.UUID, menuItem domain.MenuItem) (domain.MenuItem, error) {
	menuItemEntity, err := buildMenuItemEntity(menuIdentifier, menuItem)

	if err != nil {
		return domain.MenuItem{}, err
	}

	err = database.Create(&menuItemEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return domain.MenuItem{}, err
	}

	return converters.MenuItemEntityToMenuItemDomain(menuItemEntity), nil
}

func buildMenuEntity(menu domain.Menu) (entities.Menu, error) {
	identifier, _ := uuid.NewRandom()
	partnerEntity, err := findPartnerByIdentifier(menu.PartnerIdentifier)

	if err != nil {
		return entities.Menu{}, err
	}

	return entities.Menu{
		Identifier: identifier,
		PartnerId:  partnerEntity.Id,
	}, err
}

func findPartnerByIdentifier(identifier string) (entities.Partner, error) {
	var partnerEntity entities.Partner
	err := database.First(&partnerEntity, "partners.identifier = ?", identifier).Error

	if err != nil {
		log.Print("Find partner error: " + err.Error())

		return entities.Partner{}, err
	}

	return partnerEntity, nil
}

func buildMenuItemEntity(menuIdentifier uuid.UUID, menuItem domain.MenuItem) (entities.MenuItem, error) {
	identifier, _ := uuid.NewRandom()
	menuEntity, err := findMenuByIdentifier(menuIdentifier.String())

	if err != nil {
		return entities.MenuItem{}, err
	}

	menuItemEntity := converters.MenuItemDomainToMenuItemEntity(menuItem)
	menuItemEntity.Identifier = identifier
	menuItemEntity.MenuId = menuEntity.Id
	menuItemEntity.IsActive = true

	return menuItemEntity, nil
}

func findMenuByIdentifier(identifier string) (entities.Menu, error) {
	var menuEntity entities.Menu
	err := database.First(&menuEntity, "menus.identifier = ?", identifier).Error

	if err != nil {
		log.Print("Find menu error: " + err.Error())

		return entities.Menu{}, err
	}

	return menuEntity, nil
}

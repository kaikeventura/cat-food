package converters

import (
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

func MenuItemDomainToMenuItemEntity(menuItem domain.MenuItem) entities.MenuItem {
	return entities.MenuItem{
		Identifier:  menuItem.Identifier,
		Name:        menuItem.Name,
		Description: menuItem.Description,
		Price:       menuItem.Price,
		IsActive:    menuItem.IsActive,
	}
}

func MenuItemEntityToMenuItemDomain(menuItem entities.MenuItem) domain.MenuItem {
	return domain.MenuItem{
		Identifier:  menuItem.Identifier,
		Name:        menuItem.Name,
		Description: menuItem.Description,
		Price:       menuItem.Price,
		IsActive:    menuItem.IsActive,
	}
}

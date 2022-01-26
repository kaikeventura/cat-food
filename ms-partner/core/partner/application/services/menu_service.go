package services

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/port/outbound"
)

type MenuService struct {
	menuPersistence outbound.MenuPersistencePort
}

func ConstructMenuService(menuPersistencePort outbound.MenuPersistencePort) MenuService {
	return MenuService{
		menuPersistence: menuPersistencePort,
	}
}

func (service MenuService) CreateNewMenu(menu domain.Menu) (domain.Menu, error) {
	createdMenu, err := service.menuPersistence.SaveMenu(menu)

	if err != nil {
		return domain.Menu{}, err
	}

	return createdMenu, nil
}

func (service MenuService) CreateNewMenuItem(menuIdentifier uuid.UUID, menuItem domain.MenuItem) (domain.MenuItem, error) {
	createdMenuItem, err := service.menuPersistence.SaveMenuItem(menuIdentifier, menuItem)

	if err != nil {
		return domain.MenuItem{}, err
	}

	return createdMenuItem, nil
}

func (service MenuService) DeleteMenuItem(menuItemIdentifier uuid.UUID) error {
	err := service.menuPersistence.DeleteMenuItemByIdentifier(menuItemIdentifier)

	if err != nil {
		return err
	}

	return nil
}

func (service MenuService) ListMenuItems(menuIdentifier uuid.UUID) ([]domain.MenuItem, error) {
	menuItems, err := service.menuPersistence.ListMenuItemsByMenuIdentifier(menuIdentifier)

	if err != nil {
		return []domain.MenuItem{}, err
	}

	if len(menuItems) == 0 {
		return []domain.MenuItem{}, err
	}

	return menuItems, nil
}

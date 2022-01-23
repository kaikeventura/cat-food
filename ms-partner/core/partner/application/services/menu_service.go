package services

import (
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
	menuCreated, err := service.menuPersistence.SaveMenu(menu)

	if err != nil {
		return domain.Menu{}, err
	}

	return menuCreated, nil
}

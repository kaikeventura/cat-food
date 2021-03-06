package outbound

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

type MenuPersistencePort interface {
	SaveMenu(menu domain.Menu) (domain.Menu, error)
	SaveMenuItem(menuIdentifier uuid.UUID, menuItem domain.MenuItem) (domain.MenuItem, error)
	UpdateMenuItem(menuIdentifier uuid.UUID, menuItem domain.MenuItem) (domain.MenuItem, error)
	DeleteMenuItemByIdentifier(menuItemIdentifier uuid.UUID) error
	ListMenuItemsByMenuIdentifier(menuIdentifier uuid.UUID) ([]domain.MenuItem, error)
}

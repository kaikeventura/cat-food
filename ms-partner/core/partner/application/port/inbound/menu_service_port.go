package inbound

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

type MenuServicePort interface {
	CreateNewMenu(menu domain.Menu) (domain.Menu, error)
	CreateNewMenuItem(menuIdentifier uuid.UUID, menuItem domain.MenuItem) (domain.MenuItem, error)
	UpdateMenuItem(menuIdentifier uuid.UUID, menuItem domain.MenuItem) (domain.MenuItem, error)
	DeleteMenuItem(menuItemIdentifier uuid.UUID) error
	ListMenuItems(menuIdentifier uuid.UUID) ([]domain.MenuItem, error)
}

package outbound

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

type MenuPersistencePort interface {
	SaveMenu(menu domain.Menu) (domain.Menu, error)
	SaveMenuItem(menuIdentifier uuid.UUID, menuItem domain.MenuItem) (domain.MenuItem, error)
}

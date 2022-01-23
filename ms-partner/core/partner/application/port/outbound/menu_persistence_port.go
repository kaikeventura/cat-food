package outbound

import "github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"

type MenuPersistencePort interface {
	SaveMenu(menu domain.Menu) (domain.Menu, error)
}
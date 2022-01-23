package inbound

import "github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"

type MenuServicePort interface {
	CreateNewMenu(menu domain.Menu) (domain.Menu, error)
}

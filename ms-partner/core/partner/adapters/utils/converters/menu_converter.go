package converters

import (
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

func MenuEntityToMenuDomain(menu entities.Menu, partnerIdentifier string) domain.Menu {
	return domain.Menu{
		Identifier:        menu.Identifier,
		PartnerIdentifier: partnerIdentifier,
	}
}

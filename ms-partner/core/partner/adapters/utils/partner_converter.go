package utils

import (
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

func PartnerDomainToPartnerEntity(partnerDomain domain.Partner) entities.Partner {
	return entities.Partner{
		Identifier:     partnerDomain.Identifier,
		Name:           partnerDomain.Name,
		PartnerType:    partnerDomain.PartnerType,
		UserIdentifier: partnerDomain.UserIdentifier,
		Address:        AddressDomainToAddressEntity(partnerDomain.Address),
	}
}

func PartnerEntityToPartnerDomain(partnerEntity entities.Partner) domain.Partner {
	return domain.Partner{
		Identifier:     partnerEntity.Identifier,
		Name:           partnerEntity.Name,
		PartnerType:    partnerEntity.PartnerType,
		UserIdentifier: partnerEntity.UserIdentifier,
		Address:        AddressEntityToAddressDomain(partnerEntity.Address),
	}
}

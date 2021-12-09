package utils

import (
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

func AddressDomainToAddressEntity(addressDomain domain.Address) entities.Address {
	return entities.Address{
		Identifier: addressDomain.Identifier,
		Street:     addressDomain.Street,
		Number:     addressDomain.Number,
		District:   addressDomain.District,
		City:       addressDomain.City,
		State:      addressDomain.State,
		PostalCode: addressDomain.PostalCode,
	}
}

func AddressEntityToAddressDomain(addressEntity entities.Address) domain.Address {
	return domain.Address{
		Identifier: addressEntity.Identifier,
		Street:     addressEntity.Street,
		Number:     addressEntity.Number,
		District:   addressEntity.District,
		City:       addressEntity.City,
		State:      addressEntity.State,
		PostalCode: addressEntity.PostalCode,
	}
}

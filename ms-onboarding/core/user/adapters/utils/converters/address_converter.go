package converters

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
)

func AddressesDomainToAddressesEntity(addressesDomain []domain.Address) []entities.Address {
	var addressesEntity []entities.Address
	for _, address := range addressesDomain {
		addressesEntity = append(addressesEntity,
			entities.Address{
				Identifier: address.Identifier,
				Street:     address.Street,
				Number:     address.Number,
				District:   address.District,
				City:       address.City,
				State:      address.State,
				PostalCode: address.PostalCode,
				Main:       address.Main,
				Surname:    address.Surname,
			},
		)
	}

	return addressesEntity
}

func AddressesEntityToAddressesDomain(addressesEntity []entities.Address) []domain.Address {
	var addressesDomain []domain.Address
	for _, address := range addressesEntity {
		addressesDomain = append(addressesDomain,
			domain.Address{
				Identifier: address.Identifier,
				Street:     address.Street,
				Number:     address.Number,
				District:   address.District,
				City:       address.City,
				State:      address.State,
				PostalCode: address.PostalCode,
				Main:       address.Main,
				Surname:    address.Surname,
			},
		)
	}

	return addressesDomain
}

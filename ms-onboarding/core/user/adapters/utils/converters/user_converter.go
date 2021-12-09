package converters

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
)

func UserDomainToUserEntity(userDomain domain.User) entities.User {
	return entities.User{
		Identifier: userDomain.Identifier,
		FirstName:  userDomain.FirstName,
		LastName:   userDomain.LastName,
		BirthDate:  userDomain.BirthDate,
		Document:   userDomain.Document,
		UserType:   userDomain.UserType,
		Status: userDomain.Status,
		Addresses:  AddressesDomainToAddressesEntity(userDomain.Addresses),
	}
}

func UserEntityToUserDomain(userEntity entities.User) domain.User {
	return domain.User{
		Identifier: userEntity.Identifier,
		FirstName:  userEntity.FirstName,
		LastName:   userEntity.LastName,
		BirthDate:  userEntity.BirthDate,
		Document:   userEntity.Document,
		UserType:   userEntity.UserType,
		Status: userEntity.Status,
		Addresses:  AddressesEntityToAddressesDomain(userEntity.Addresses),
	}
}

package persistence

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

type MySQLUserPersistence struct {
}

func ConstructMySQLUserPersistence(databaseRepository *gorm.DB) *MySQLUserPersistence {
	database = databaseRepository

	return &MySQLUserPersistence{}
}

func (MySQLUserPersistence) SaveUser(user domain.UserDomain) (domain.UserDomain, error) {
	addressesEntity := buildAddressesEntity(user.Address)
	userEntity := buildUserEntity(user, addressesEntity)
	err := database.Create(&userEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())
	}

	return user, err
}

func buildAddressesEntity(addresses []domain.Address) []entities.Address {
	if !checkForAddresses(addresses) {
		return []entities.Address{}
	}

	return addressesDomainToAddressesEntity(addresses)
}

func checkForAddresses(addresses []domain.Address) bool {
	return len(addresses) > 0 || addresses != nil
}

func addressesDomainToAddressesEntity(addressesDomain []domain.Address) []entities.Address {
	var addressesEntity []entities.Address
	for _, address := range addressesDomain {
		addressesEntity = append(addressesEntity,
			entities.Address{
				Street:   address.Street,
				Number:   address.Number,
				District: address.District,
				City:     address.City,
				State:    address.State,
			},
		)
	}

	return addressesEntity
}

func buildUserEntity(userDomain domain.UserDomain, addressesEntity []entities.Address) entities.User {
	userEntity := userDomainToUserEntity(userDomain)
	userEntity.Address = addressesEntity

	return userEntity
}

func userDomainToUserEntity(userDomain domain.UserDomain) entities.User {
	return entities.User{
		FirstName: userDomain.FirstName,
		LastName:  userDomain.LastName,
		BirthDate: userDomain.BirthDate,
		Document:  userDomain.Document,
		UserType:  userDomain.UserType,
	}
}

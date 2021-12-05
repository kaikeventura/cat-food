package persistence

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/utils/converters"
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

func (MySQLUserPersistence) SaveUser(user domain.User) (domain.User, error) {
	addressesEntity := buildAddressesEntity(user.Address)
	userEntity := buildUserEntity(user, addressesEntity)
	err := database.Create(&userEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())
	}

	return converters.UserEntityToUserDomain(userEntity), err
}

func buildAddressesEntity(addresses []domain.Address) []entities.Address {
	if !checkForAddresses(addresses) {
		return []entities.Address{}
	}

	addressesEntity := converters.AddressesDomainToAddressesEntity(addresses)
	addRandomUUIDToAddresses(addressesEntity)

	return addressesEntity
}

func checkForAddresses(addresses []domain.Address) bool {
	return len(addresses) > 0 || addresses != nil
}

func addRandomUUIDToAddresses(addressesEntity []entities.Address) {
	for i := 0; i < len(addressesEntity); i++ {
		addressesEntity[i].Identifier, _ = uuid.NewRandom()
	}
}

func buildUserEntity(userDomain domain.User, addressesEntity []entities.Address) entities.User {
	userEntity := converters.UserDomainToUserEntity(userDomain)
	userEntity.Identifier, _ = uuid.NewRandom()
	userEntity.Address = addressesEntity

	return userEntity
}

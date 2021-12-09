package persistence

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/utils/converters"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/sub_domain"
	"gorm.io/gorm"
	"log"
)

type MySQLUserPersistence struct {
	database *gorm.DB
}

func ConstructMySQLUserPersistence(databaseRepository *gorm.DB) MySQLUserPersistence {
	return MySQLUserPersistence{
		database: databaseRepository,
	}
}

func (persistence MySQLUserPersistence) SaveUser(user domain.User) (domain.User, error) {
	addressesEntity := buildAddressesEntity(user.Addresses)
	userEntity := buildUserEntity(user, addressesEntity)
	err := persistence.database.Create(&userEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return domain.User{}, err
	}

	return converters.UserEntityToUserDomain(userEntity), err
}

func (persistence MySQLUserPersistence) FindUserByIdentifier(identifier uuid.UUID) (domain.User, error) {
	var user entities.User
	err := persistence.database.Preload("Addresses").Joins("INNER JOIN addresses ON addresses.user_id = users.id").Where("users.identifier = ?", identifier).First(&user).Error

	if err != nil {
		log.Print("Find user error: " + err.Error())

		return domain.User{}, err
	}

	return converters.UserEntityToUserDomain(user), err
}

func (persistence MySQLUserPersistence) CheckUserStatus(identifier uuid.UUID) (sub_domain.UserStatus, error) {
	var user entities.User
	err := persistence.database.Select("users.status").Where("users.identifier = ?", identifier).First(&user).Error

	if err != nil {
		log.Print("Find user error: " + err.Error())

		return sub_domain.UserStatus{}, err
	}

	return sub_domain.UserStatus{Status: user.Status}, nil
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
	userEntity.Status = "ENABLED"
	userEntity.Addresses = addressesEntity

	return userEntity
}

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
	address := []entities.Address {
		{
			Street: user.Address[0].Street,
			Number: user.Address[0].Number,
			District: user.Address[0].District,
			City: user.Address[0].City,
			State: user.Address[0].State,
		},
	}

	userEntity := entities.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		BirthDate: user.BirthDate,
		Document: user.Document,
		UserType: user.UserType,
		Address: address,
	}
	err := database.Create(&userEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())
	}

	return user, err
}

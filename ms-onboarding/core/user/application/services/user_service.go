package services

import (
	"errors"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/port/outbound"
)

var userPersistence outbound.UserPersistencePort

type UserService struct {
}

func ConstructUserService(persistence outbound.UserPersistencePort) *UserService {
	userPersistence = persistence

	return &UserService{}
}

func (UserService) CreateNewUser(user domain.User) (domain.User, error) {
	err := checkIfASingleAddressIsThePrimary(user.Address)
	if err != nil {
		return user, err
	}

	return userPersistence.SaveUser(user)
}

func checkIfASingleAddressIsThePrimary(addresses []domain.Address) error {
	var mainAddresses int8
	for _, address := range addresses {
		if address.Main {
			mainAddresses++
		}
	}

	if mainAddresses > 1 {
		return errors.New("Only one main address is allowed")
	}

	return nil
}

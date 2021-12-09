package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/port/outbound"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/sub_domain"
)

type UserService struct {
	userPersistence outbound.UserPersistencePort
}

func ConstructUserService(persistence outbound.UserPersistencePort) UserService {
	return UserService{
		userPersistence: persistence,
	}
}

func (service UserService) CreateNewUser(user domain.User) (domain.User, error) {
	err := checkIfASingleAddressIsThePrimary(user.Addresses)
	if err != nil {
		return user, err
	}

	return service.userPersistence.SaveUser(user)
}

func (service UserService) FindUser(identifier uuid.UUID) (domain.User, error) {
	return service.userPersistence.FindUserByIdentifier(identifier)
}

func (service UserService) CheckUserStatus(identifier uuid.UUID) (sub_domain.UserStatus, error) {
	return service.userPersistence.CheckUserStatus(identifier)
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

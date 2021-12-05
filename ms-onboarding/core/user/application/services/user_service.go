package services

import (
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

	return userPersistence.SaveUser(user)
}

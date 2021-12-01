package services

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
)

type UserService struct {
}

func ConstructUserService() *UserService {
	return &UserService{}
}

func (UserService) CreateNewUser(user domain.UserDomain) domain.UserDomain {
	return user
}

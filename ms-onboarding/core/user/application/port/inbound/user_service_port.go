package inbound

import "github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"

type UserServicePort interface {
	CreateNewUser(user domain.UserDomain) (domain.UserDomain, error)
}

package inbound

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/sub_domain"
)

type UserServicePort interface {
	CreateNewUser(user domain.User) (domain.User, error)
	FindUser(identifier uuid.UUID) (domain.User, error)
	UserDetails(identifier uuid.UUID) (sub_domain.UserStatus, error)
}

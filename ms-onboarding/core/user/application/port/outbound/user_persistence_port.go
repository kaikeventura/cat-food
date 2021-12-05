package outbound

import "github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"

type UserPersistencePort interface {
	SaveUser(user domain.User) (domain.User, error)
}

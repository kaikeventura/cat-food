package outbound

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/sub_domain"
)

type UserPersistencePort interface {
	SaveUser(user domain.User) (domain.User, error)
	FindUserByIdentifier(identifier uuid.UUID) (domain.User, error)
	GetUserDetails(identifier uuid.UUID) (sub_domain.UserStatus, error)
}

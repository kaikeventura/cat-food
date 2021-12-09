package outbound

import "github.com/kaikeventura/cat-food/ms-partner/core/partner/application/sub_domain"

type OnboardingClientPort interface {
	CheckIfUserExists(userIdentifier string) (sub_domain.UserStatusResponse, error)
}

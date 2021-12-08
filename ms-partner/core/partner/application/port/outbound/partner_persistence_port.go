package outbound

import "github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"

type PartnerPersistencePort interface {
	SavePartner(partner domain.Partner) (domain.Partner, error)
}

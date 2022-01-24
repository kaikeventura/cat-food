package outbound

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

type PartnerPersistencePort interface {
	SavePartner(partner domain.Partner) (domain.Partner, error)
	FindByIdentifier(identifier uuid.UUID) (domain.Partner, error)
}

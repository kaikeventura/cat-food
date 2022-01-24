package inbound

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
)

type PartnerServicePort interface {
	CreateNewPartner(partner domain.Partner) (domain.Partner, error)
	GetPartner(identifier uuid.UUID) (domain.Partner, error)
}

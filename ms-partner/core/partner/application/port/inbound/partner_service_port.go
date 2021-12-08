package inbound

import "github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"

type PartnerServicePort interface {
	CreateNewPartner(partner domain.Partner) (domain.Partner, error)
}

package services

import (
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/port/outbound"
)

type PartnerService struct {
	partnerPersistence outbound.PartnerPersistencePort
}

func ConstructPartnerService(partnerPersistencePort outbound.PartnerPersistencePort) PartnerService {
	return PartnerService{
		partnerPersistence: partnerPersistencePort,
	}
}

func (service PartnerService) CreateNewPartner(partner domain.Partner) (domain.Partner, error) {
	createdPartner, err := service.partnerPersistence.SavePartner(partner)
	if err != nil {
		return domain.Partner{}, err
	}

	return createdPartner, err
}

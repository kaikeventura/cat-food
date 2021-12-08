package services

import "github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"

type PartnerService struct {
}

func ConstructPartnerService() PartnerService {
	return PartnerService{}
}

func (service PartnerService) CreateNewPartner(partner domain.Partner) (domain.Partner, error) {
	return partner, nil
}

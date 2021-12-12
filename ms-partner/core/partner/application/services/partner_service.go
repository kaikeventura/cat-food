package services

import (
	"errors"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/port/outbound"
)

type PartnerService struct {
	partnerPersistence outbound.PartnerPersistencePort
	onboardingClient   outbound.OnboardingClientPort
}

func ConstructPartnerService(partnerPersistencePort outbound.PartnerPersistencePort, onboardingClientPort outbound.OnboardingClientPort) PartnerService {
	return PartnerService{
		partnerPersistence: partnerPersistencePort,
		onboardingClient:   onboardingClientPort,
	}
}

func (service PartnerService) CreateNewPartner(partner domain.Partner) (domain.Partner, error) {
	userStatus, err := service.onboardingClient.CheckIfUserExists(partner.UserIdentifier)

	if err != nil {
		return domain.Partner{}, err
	}

	if userStatus.Status != "ENABLED" {
		return domain.Partner{}, errors.New("User not enabled ")
	}

	createdPartner, err := service.partnerPersistence.SavePartner(partner)
	if err != nil {
		return domain.Partner{}, err
	}

	return createdPartner, err
}

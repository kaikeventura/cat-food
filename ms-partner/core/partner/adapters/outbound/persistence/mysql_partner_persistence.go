package persistence

import (
	"log"

	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/utils/converters"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
	"gorm.io/gorm"
)

type MySQLPartnerPersistence struct {
	database *gorm.DB
}

func ConstructMySQLPartnerPersistence(databaseRepository *gorm.DB) MySQLPartnerPersistence {
	return MySQLPartnerPersistence{
		database: databaseRepository,
	}
}

func (persistence MySQLPartnerPersistence) SavePartner(partner domain.Partner) (domain.Partner, error) {
	partnerEntity := buildPartnerEntity(partner)
	err := persistence.database.Create(&partnerEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return domain.Partner{}, err
	}

	return converters.PartnerEntityToPartnerDomain(partnerEntity), err
}

func (persistence MySQLPartnerPersistence) FindByIdentifier(identifier uuid.UUID) (domain.Partner, error) {
	var partner entities.Partner
	err := persistence.database.
		Preload("Address").
		Joins("INNER JOIN addresses ON addresses.partner_id = partners.id").
		Preload("Menu").
		Joins("INNER JOIN menus ON menus.partner_id = partners.id").
		Where("partners.identifier = ?", identifier).
		First(&partner).Error

	if err != nil {
		log.Print("Find partner error: " + err.Error())

		return domain.Partner{}, err
	}

	return converters.PartnerEntityToPartnerDomain(partner), err
}

func buildPartnerEntity(partner domain.Partner) entities.Partner {
	partnerEntity := converters.PartnerDomainToPartnerEntity(partner)
	partnerEntity.Identifier, _ = uuid.NewRandom()
	partnerEntity.Address.Identifier, _ = uuid.NewRandom()

	return partnerEntity
}

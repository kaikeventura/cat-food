package persistence

import (
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence/entities"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/utils"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
	"gorm.io/gorm"
	"log"
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
	partnerEntity := buildUserEntity(partner)
	err := persistence.database.Create(&partnerEntity).Error

	if err != nil {
		log.Print("Persistence error: " + err.Error())

		return domain.Partner{}, err
	}

	return utils.PartnerEntityToPartnerDomain(partnerEntity), err
}

func buildUserEntity(partner domain.Partner) entities.Partner {
	partnerEntity := utils.PartnerDomainToPartnerEntity(partner)
	partnerEntity.Identifier, _ = uuid.NewRandom()
	partnerEntity.Address.Identifier, _ = uuid.NewRandom()

	return partnerEntity
}

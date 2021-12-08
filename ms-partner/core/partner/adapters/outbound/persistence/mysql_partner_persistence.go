package persistence

import (
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

	panic("implement me")
}

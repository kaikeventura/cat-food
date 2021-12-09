package container

import (
	"github.com/kaikeventura/cat-food/ms-partner/configuration/database"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/inbound/controllers"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/http"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/services"
)

func RunDependencyInjection() {
	mySQLPartnerPersistence := persistence.ConstructMySQLPartnerPersistence(database.GetDatabase())
	constructOnboardingClient := http.ConstructOnboardingClient()
	partnerService := services.ConstructPartnerService(mySQLPartnerPersistence, constructOnboardingClient)
	controllers.ConstructUserController(partnerService)
}
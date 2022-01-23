package container

import (
	"github.com/kaikeventura/cat-food/ms-partner/configuration/database"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/inbound/controllers"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/http"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/outbound/persistence"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/services"
)

func RunDependencyInjection() {
	partnerDependencyInjection()
	menuDependencyInjection()
}

func partnerDependencyInjection() {
	mySQLPartnerPersistence := persistence.ConstructMySQLPartnerPersistence(database.GetDatabase())
	onboardingClient := http.ConstructOnboardingClient()
	partnerService := services.ConstructPartnerService(mySQLPartnerPersistence, onboardingClient)
	controllers.ConstructUserController(partnerService)
}

func menuDependencyInjection() {
	mySQLMenuPersistence := persistence.ConstructMySQLMenuPersistence(database.GetDatabase())
	menuService := services.ConstructMenuService(mySQLMenuPersistence)
	controllers.ConstructMenuController(menuService)
}

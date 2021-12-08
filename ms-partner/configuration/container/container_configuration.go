package container

import (
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/inbound/controllers"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/services"
)

func RunDependencyInjection() {
	partnerService := services.ConstructPartnerService()
	controllers.ConstructUserController(partnerService)
}

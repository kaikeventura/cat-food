package container

import (
	"github.com/kaikeventura/cat-food/ms-onboarding/configuration/database"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/inbound/controllers"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/outbound/persistence"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/services"
)

func RunDependencyInjection() {
	mySQLPersistence := persistence.ConstructMySQLUserPersistence(database.GetDatabase())
	userService := services.ConstructUserService(mySQLPersistence)
	controllers.ConstructUserController(userService)
}

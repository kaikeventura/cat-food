package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/adapters/inbound/controllers"
)

func ConfigurationRouter(router *gin.Engine) *gin.Engine {
	main := router.Group("v1/")
	{
		users := main.Group("users")
		{
			users.POST("/", controllers.CreateUser)
			users.GET("/:identifier", controllers.GetUser)
		}
	}

	return router
}

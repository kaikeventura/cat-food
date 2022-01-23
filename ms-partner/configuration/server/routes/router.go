package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/adapters/inbound/controllers"
)

func ConfigurationRouter(router *gin.Engine) *gin.Engine {
	main := router.Group("v1/")
	{
		partners := main.Group("partners")
		{
			partners.POST("/", controllers.CreatePartner)
		}

		menus := main.Group("menus")
		{
			menus.POST("/", controllers.CreateMenu)
		}
	}

	return router
}

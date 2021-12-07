package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-food/ms-partners/core/partner/adapters/inbound/controllers"
)

func ConfigurationRouter(router *gin.Engine) *gin.Engine {
	main := router.Group("v1/")
	{
		users := main.Group("partners")
		{
			users.POST("/", controllers.CreatePartner)
		}
	}

	return router
}

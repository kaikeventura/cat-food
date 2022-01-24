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
			partners.GET("/:user_identifier", controllers.GetPartner)
		}

		menus := main.Group("menus")
		{
			menus.POST("/", controllers.CreateMenu)
			menus.POST("/:menu_identifier/items/add", controllers.AddItem)
			menus.DELETE("/items/:item_identifier/remove", controllers.RemoveItem)
			menus.GET("/items/:menu_identifier", controllers.ListItems)
		}
	}

	return router
}

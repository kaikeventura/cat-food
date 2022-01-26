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
			menus.POST("/:menu_identifier/items", controllers.AddMenuItem)
			menus.PUT("/items/:menu_item_identifier", controllers.UpdateMenuItem)
			menus.DELETE("/items/:menu_item_identifier", controllers.RemoveMenuItem)
			menus.GET("/items/:menu_identifier", controllers.ListMenuItems)
		}
	}

	return router
}

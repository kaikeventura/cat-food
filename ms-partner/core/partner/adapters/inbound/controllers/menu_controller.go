package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/port/inbound"
)

var menuService inbound.MenuServicePort

func ConstructMenuController(service inbound.MenuServicePort) {
	menuService = service
}

func CreateMenu(context *gin.Context) {
	var menu domain.Menu
	err := context.ShouldBindJSON(&menu)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	newMenu, err := menuService.CreateNewMenu(menu)

	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(201, newMenu)
}

func AddItem(context *gin.Context) {

}

func RemoveItem(context *gin.Context) {

}

func ListItems(context *gin.Context) {

}

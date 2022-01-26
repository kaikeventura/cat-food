package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func AddMenuItem(context *gin.Context) {
	menuIdentifier := context.Param("menu_identifier")
	identifierUUID, err := uuid.Parse(menuIdentifier)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Identifier has to be UUID",
		})

		return
	}

	var menuItem domain.MenuItem
	jsonBindErr := context.ShouldBindJSON(&menuItem)

	if jsonBindErr != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	createdMenuItem, err := menuService.CreateNewMenuItem(identifierUUID, menuItem)

	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(201, createdMenuItem)
}

func UpdateMenuItem(context *gin.Context) {

}

func RemoveMenuItem(context *gin.Context) {

}

func ListMenuItems(context *gin.Context) {

}

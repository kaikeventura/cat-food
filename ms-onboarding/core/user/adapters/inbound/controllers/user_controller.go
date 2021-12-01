package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/port/inbound"
)

var userService inbound.UserServicePort

func ConstructUserController(service inbound.UserServicePort)  {
	userService = service
}

func CreateUser(context *gin.Context) {
	var user domain.UserDomain
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	var result = userService.CreateNewUser(user)

	context.JSON(200, gin.H{
		"success": result,
	})
}

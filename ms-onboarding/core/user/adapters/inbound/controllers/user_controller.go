package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/domain"
	"github.com/kaikeventura/cat-food/ms-onboarding/core/user/application/port/inbound"
)

var userService inbound.UserServicePort

func ConstructUserController(service inbound.UserServicePort) {
	userService = service
}

func CreateUser(context *gin.Context) {
	var user domain.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	var newUser, createUserError = userService.CreateNewUser(user)

	if createUserError != nil {
		context.JSON(404, gin.H{
			"error": createUserError.Error(),
		})

		return
	}

	context.JSON(200, newUser)
}

func GetUser(context *gin.Context) {
	identifier := context.Param("identifier")
	identifierUUID, err := uuid.Parse(identifier)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Identifier has to be UUID",
		})

		return
	}
	
	user, err := userService.FindUser(identifierUUID)

	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(200, user)
}

func UserStatus(context *gin.Context) {
	identifier := context.Param("identifier")
	identifierUUID, err := uuid.Parse(identifier)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Identifier has to be UUID",
		})

		return
	}

	userStatus, err := userService.CheckUserStatus(identifierUUID)
	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(200, userStatus)
}

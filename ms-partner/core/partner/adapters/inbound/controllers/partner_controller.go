package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/domain"
	"github.com/kaikeventura/cat-food/ms-partner/core/partner/application/port/inbound"
)

var partnerService inbound.PartnerServicePort

func ConstructUserController(service inbound.PartnerServicePort) {
	partnerService = service
}

func CreatePartner(context *gin.Context) {
	var partner domain.Partner
	err := context.ShouldBindJSON(&partner)

	if err != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind Json: " + err.Error(),
		})

		return
	}

	newPartner, err := partnerService.CreateNewPartner(partner)

	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	context.JSON(201, newPartner)
}

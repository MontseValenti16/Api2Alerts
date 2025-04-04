package handlers

import (
	"LifeGuardAlertas/src/samrtwatch/application"
	"LifeGuardAlertas/src/samrtwatch/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateSmartwatchController struct {
	updateSmartwatchUseCase *application.UpdateSmartwatchUseCase
}

func NewUpdateSmartwatchController(updateSmartwatchUseCase *application.UpdateSmartwatchUseCase) *UpdateSmartwatchController {
	return &UpdateSmartwatchController{
		updateSmartwatchUseCase: updateSmartwatchUseCase,
	}
}

func (c *UpdateSmartwatchController) UpdateByMacAddress(ctx *gin.Context) {
	macAddress := ctx.Param("macaddress")

	var data entities.Smartwatch
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.updateSmartwatchUseCase.Run(macAddress, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Smartwatch actualizada correctamente",
	})
}
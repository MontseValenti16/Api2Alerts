package handlers

import (
	"LifeGuardAlertas/src/samrtwatch/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetSmartwatchController struct {
	getSmartwatchUseCase *application.GetSmartwatchUseCase
}

func NewGetSmartwatchController(getSmartwatchUseCase *application.GetSmartwatchUseCase) *GetSmartwatchController {
	return &GetSmartwatchController{
		getSmartwatchUseCase: getSmartwatchUseCase,
	}
}

func (c *GetSmartwatchController) GetByMacAddress(ctx *gin.Context) {
	macAddress := ctx.Param("macaddress")

	smartwatch, err := c.getSmartwatchUseCase.Run(macAddress)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Smartwatch no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, smartwatch)
}

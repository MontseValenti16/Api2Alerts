// infrastructure/handlers/Alerta_Controller.go
package handlers

import (
	"LifeGuardAlertas/src/alerta/application"
	"LifeGuardAlertas/src/alerta/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAlertaController struct {
	getAlertasUseCase   *application.GetAlertasUseCase
}

func NewGetAlertaController(
	getAlertasUseCase *application.GetAlertasUseCase,
) *GetAlertaController {
	return &GetAlertaController{
		getAlertasUseCase:   getAlertasUseCase,
	}
}


func (c *GetAlertaController) GetAllAlertas(ctx *gin.Context) {
	alertas, err := c.getAlertasUseCase.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, alertas)
}

func (c *GetAlertaController) GetAlertasByMacAddress(ctx *gin.Context) {
	macAddress := ctx.Param("macaddress")
	if macAddress == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "macaddress es requerido"})
		return
	}

	alertas, err := c.getAlertasUseCase.RunByMacAddress(macAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, alertas)
}
package handlers

import (
	"LifeGuardAlertas/src/alerta/application"
	"LifeGuardAlertas/src/alerta/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAlertaController struct {
	createAlertaUseCase *application.CreateAlertaUseCase
}

func NewCreateAlertaController(createAlertaUseCase *application.CreateAlertaUseCase) *CreateAlertaController {
	return &CreateAlertaController{
		createAlertaUseCase: createAlertaUseCase,
	}
}

func (c *CreateAlertaController) SaveFiebre(ctx *gin.Context) {
	var data entities.Alerta
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Guardado"})
}

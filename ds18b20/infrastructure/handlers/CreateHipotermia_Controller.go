package handlers

import (
	"LifeGuardAlertas/ds18b20/application"
	"LifeGuardAlertas/ds18b20/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateHipotermiaController struct {
	createHipotermiaUseCae *application.CreateHipotermiaUseCase
}

func NewCreateHipotermiaController(creaHipotermiaUseCase *application.CreateHipotermiaUseCase) *CreateHipotermiaController{
	return &CreateHipotermiaController{
		createHipotermiaUseCae: creaHipotermiaUseCase,
	}
}

func (c *CreateHipotermiaController) SaveHipotermia(ctx *gin.Context) {
	var data *entities.AlertaDs18b20
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Guardado"})
}
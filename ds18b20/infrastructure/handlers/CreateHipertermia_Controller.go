package handlers

import (
	"LifeGuardAlertas/ds18b20/application"
	"LifeGuardAlertas/ds18b20/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateHipertermiaController struct {
	createHipertermiaUseCase *application.CreateHipertermiaUseCase
}

func NewCreateHipertermiaController(creaHipertermiaUseCase *application.CreateHipertermiaUseCase) *CreateHipertermiaController{
	return &CreateHipertermiaController{
		createHipertermiaUseCase: creaHipertermiaUseCase,
	}
}

func (c *CreateHipertermiaController) SaveHipertermia(ctx *gin.Context) {
	var data *entities.AlertaDs18b20
	if err := ctx.ShouldBindJSON(&data); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	if err := ctx.ShouldBindJSON(data); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	ctx.JSON(http.StatusOK, gin.H{"message": "Guardado"})
}

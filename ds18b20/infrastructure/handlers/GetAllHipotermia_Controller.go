package handlers

import (
	"LifeGuardAlertas/ds18b20/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllHipotermiaController struct {
	getAllHipotermiaController *application.GetAllHipotermiaUseCase
}

func NewGetAllHipotermiaController(getAllHipotermiaController *application.GetAllHipotermiaUseCase) *GetAllHipotermiaController{
	return &GetAllHipotermiaController{
		getAllHipotermiaController: getAllHipotermiaController,
	}
}

func (c *GetAllHipotermiaController) GetAllHipotermia(ctx *gin.Context) {
	data, err := c.getAllHipotermiaController.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}
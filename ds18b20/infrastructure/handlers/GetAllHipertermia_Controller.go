package handlers

import (
	"LifeGuardAlertas/ds18b20/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllHipertermiaController struct {
	getAllHipertermiaController *application.GetAllHipertermiaUseCase
}

func NewGetAllHipertermiaController(getAllHipertermiaController *application.GetAllHipertermiaUseCase) *GetAllHipertermiaController{
	return &GetAllHipertermiaController{
		getAllHipertermiaController: getAllHipertermiaController,
	}
}

func (c *GetAllHipertermiaController) GetAllHipertermia(ctx *gin.Context) {
	data, err := c.getAllHipertermiaController.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK , gin.H{"data": data})
}
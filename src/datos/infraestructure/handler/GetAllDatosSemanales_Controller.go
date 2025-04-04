package handlers

import (
	"LifeGuardAlertas/src/datos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDatosSemanalesController struct {
	useCase *application.GetAllDatosSemanalUseCase
}

func NewGetAllDatosSemanalesController(useCase *application.GetAllDatosSemanalUseCase) *GetAllDatosSemanalesController {
	return &GetAllDatosSemanalesController{
		useCase: useCase,
	}
}

func (c *GetAllDatosSemanalesController) Handle(ctx *gin.Context) {
	datos, err := c.useCase.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": datos,
	})
}

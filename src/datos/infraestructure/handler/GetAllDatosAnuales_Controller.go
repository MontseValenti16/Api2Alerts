package handlers

import (
	"LifeGuardAlertas/src/datos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDatosAnualesController struct {
	useCase *application.GetAllDatosAnualesUseCase
}

func NewGetAllDatosAnualesController(useCase *application.GetAllDatosAnualesUseCase) *GetAllDatosAnualesController {
	return &GetAllDatosAnualesController{
		useCase: useCase,
	}
}

func (c *GetAllDatosAnualesController) Handle(ctx *gin.Context) {
	datos, err := c.useCase.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": datos,
	})
}

package handlers

import (
	"LifeGuardAlertas/src/datos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDatosDiariosController struct {
	useCase *application.GetAllDatosDiariosUseCase
}

func NewGetAllDatosDiariosController(useCase *application.GetAllDatosDiariosUseCase) *GetAllDatosDiariosController {
	return &GetAllDatosDiariosController{
		useCase: useCase,
	}
}

func (c *GetAllDatosDiariosController) Handle(ctx *gin.Context) {
	datos, err := c.useCase.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": datos,
	})
}

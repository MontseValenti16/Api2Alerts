package handlers

import (
	"LifeGuardAlertas/src/datos/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDatosMensualesController struct {
    useCase *application.GetAllDatosMensualesUseCase
}

func NewGetAllDatosMensualesController(useCase *application.GetAllDatosMensualesUseCase) *GetAllDatosMensualesController {
    return &GetAllDatosMensualesController{
        useCase: useCase,
    }
}

func (c *GetAllDatosMensualesController) Handle(ctx *gin.Context) {
    datos, err := c.useCase.Run()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "data": datos,
    })
}
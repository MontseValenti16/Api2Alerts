package handlers

import (
	"LifeGuardAlertas/ds18b20/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllFiebreController struct {
	getAllFiebreController *application.GetAllFiebreUseCase
}

func NewGetAllFiebreController(getAllFiebreController *application.GetAllFiebreUseCase) *GetAllFiebreController{
	return &GetAllFiebreController{
		getAllFiebreController: getAllFiebreController,
	}
}

func (c *GetAllFiebreController) GetAllFiebre(ctx *gin.Context) {
	data, err := c.getAllFiebreController.Run()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": data})
}
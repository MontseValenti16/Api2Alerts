package controller

import (
	"LifeGuardAlertas/src/alerta/application"
	"LifeGuardAlertas/src/alerta/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type AlertaController struct {
	createAlertaUseCase *handlers.CreateAlertaController
}

func NewAlertaController(
	createUseCase *application.CreateAlertaUseCase,
) *AlertaController{
	createHandler := handlers.NewCreateAlertaController(createUseCase)
	return &AlertaController{
		createAlertaUseCase: createHandler,
	}
}

func (cc *AlertaController) CreateFiebre(ctx *gin.Context){
	cc.createAlertaUseCase.SaveFiebre(ctx)
}

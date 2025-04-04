package controller

import (
	"LifeGuardAlertas/src/alerta/application"
	"LifeGuardAlertas/src/alerta/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type AlertaController struct {
	createAlertaUseCase *handlers.CreateAlertaController
	getAlertaUseCase *handlers.GetAlertaController
}

func NewAlertaController(
	createUseCase *application.CreateAlertaUseCase,
	getUseCase *application.GetAlertasUseCase,
) *AlertaController{
	createHandler := handlers.NewCreateAlertaController(createUseCase)
	getHandler := handlers.NewGetAlertaController(getUseCase)
	return &AlertaController{
		createAlertaUseCase: createHandler,
		getAlertaUseCase: getHandler,
	}
}

func (cc *AlertaController) CreateFiebre(ctx *gin.Context){
	cc.createAlertaUseCase.SaveFiebre(ctx)
}


func (cc *AlertaController) GetAllAlertas(ctx *gin.Context) {
	cc.getAlertaUseCase.GetAllAlertas(ctx)
}

func (cc *AlertaController) GetAlertasByMacAddress(ctx *gin.Context) {
	cc.getAlertaUseCase.GetAlertasByMacAddress(ctx)
}

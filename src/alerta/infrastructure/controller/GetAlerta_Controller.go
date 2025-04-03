// infrastructure/controller/Alerta_Controller.go
package controller

import (
	"LifeGuardAlertas/src/alerta/application"
	"LifeGuardAlertas/src/alerta/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type GetAlertaController struct {
	getAlertaUseCase *handlers.GetAlertaController
}

func NewGetAlertaController(	
	getUseCase *application.GetAlertasUseCase,
) *GetAlertaController {
	getHandler := handlers.NewGetAlertaController(getUseCase)
	return &GetAlertaController{
		getAlertaUseCase: getHandler,
	}
}



func (cc *GetAlertaController) GetAllAlertas(ctx *gin.Context) {
	cc.getAlertaUseCase.GetAllAlertas(ctx)
}

func (cc *GetAlertaController) GetAlertasByMacAddress(ctx *gin.Context) {
	cc.getAlertaUseCase.GetAlertasByMacAddress(ctx)
}
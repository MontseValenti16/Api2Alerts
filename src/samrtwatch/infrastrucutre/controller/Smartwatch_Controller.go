package controller

import (
	"LifeGuardAlertas/src/samrtwatch/application"
	"LifeGuardAlertas/src/samrtwatch/infrastrucutre/handlers"

	"github.com/gin-gonic/gin"
)

type SmartwatchController struct {
	createHandler *handlers.CreateSmartwatchController
	getHandler    *handlers.GetSmartwatchController
	updateHandler *handlers.UpdateSmartwatchController
}

func NewSmartwatchController(
	createUseCase *application.CreateSmartwatchUseCase,
	getUseCase *application.GetSmartwatchUseCase,
	updateUseCase *application.UpdateSmartwatchUseCase,
) *SmartwatchController {
	return &SmartwatchController{
		createHandler: handlers.NewCreateSmartwatchController(createUseCase),
		getHandler:    handlers.NewGetSmartwatchController(getUseCase),
		updateHandler: handlers.NewUpdateSmartwatchController(updateUseCase),
	}
}

func (cc *SmartwatchController) CreateSmartwatch(ctx *gin.Context) {
	cc.createHandler.SaveSmartwatch(ctx)
}

func (cc *SmartwatchController) GetSmartwatch(ctx *gin.Context) {
	cc.getHandler.GetByMacAddress(ctx)
}

func (cc *SmartwatchController) UpdateSmartwatch(ctx *gin.Context) {
	cc.updateHandler.UpdateByMacAddress(ctx)
}
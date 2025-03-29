package controller

import (
	"LifeGuardAlertas/ds18b20/application"
	"LifeGuardAlertas/ds18b20/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type HipotermiaController struct {
	createHipotermiaUseCase *handlers.CreateHipotermiaController
	getHipotermiaUseCase *handlers.GetAllHipotermiaController
}

func NewHipotermiaController(
	createHipotermiaUseCase *application.CreateHipotermiaUseCase,
	getHipotermiaUseCase *application.GetAllHipotermiaUseCase,
) *HipotermiaController{
	createHanlder := handlers.NewCreateHipotermiaController(createHipotermiaUseCase)
	getHandler := handlers.NewGetAllHipotermiaController(getHipotermiaUseCase)
	return &HipotermiaController{
		createHipotermiaUseCase: createHanlder,
		getHipotermiaUseCase: getHandler,
	}
}

func (cc *HipotermiaController) CreateHipotermia(ctx *gin.Context){
	cc.createHipotermiaUseCase.SaveHipotermia(ctx)
}

func (gc *HipotermiaController) GetHipotermia(ctx *gin.Context){
	gc.getHipotermiaUseCase.GetAllHipotermia(ctx)
}
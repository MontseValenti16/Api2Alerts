package controller

import (
	"LifeGuardAlertas/ds18b20/application"
	"LifeGuardAlertas/ds18b20/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type HipertermiaController struct {
	createHipertermiaUseCase *handlers.CreateHipertermiaController
	getHipertermiaUseCase *handlers.GetAllHipertermiaController
}

func NewHipertermiaController(
	createUseCase *application.CreateHipertermiaUseCase,
	getUseCase *application.GetAllHipertermiaUseCase,
) *HipertermiaController {
	createHandler := handlers.NewCreateHipertermiaController(createUseCase)
	getHandler := handlers.NewGetAllHipertermiaController(getUseCase)
	return &HipertermiaController{
		createHipertermiaUseCase: createHandler,
		getHipertermiaUseCase: getHandler,
	}
}

func (cc *HipertermiaController) CreateHipertermia(ctx *gin.Context){
	cc.createHipertermiaUseCase.SaveHipertermia(ctx)
}

func (gc *HipertermiaController) GetHipertermia(ctx *gin.Context){
	gc.getHipertermiaUseCase.GetAllHipertermia(ctx)
}
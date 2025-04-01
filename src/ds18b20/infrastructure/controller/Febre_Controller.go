package controller

import (
	"LifeGuardAlertas/ds18b20/application"
	"LifeGuardAlertas/ds18b20/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

type FiebreController struct {
	createFiebreUseCase *handlers.CreateFiebreController
	getFiebreUseCase *handlers.GetAllFiebreController
}

func NewFiebreController(
	createUseCase *application.CreateFiebreUseCase,
	getUseCase *application.GetAllFiebreUseCase,
) *FiebreController{
	createHandler := handlers.NewCreateFiebreController(createUseCase)
	getHandler := handlers.NewGetAllFiebreController(getUseCase)
	return &FiebreController{
		createFiebreUseCase: createHandler,
		getFiebreUseCase: getHandler,
	}
}

func (cc *FiebreController) CreateFiebre(ctx *gin.Context){
	cc.createFiebreUseCase.SaveFiebre(ctx)
}

func (gc *FiebreController) GetFiebre(ctx *gin.Context){
	gc.getFiebreUseCase.GetAllFiebre(ctx)
}
	package handlers

	import (
		"LifeGuardAlertas/src/samrtwatch/application"
		"LifeGuardAlertas/src/samrtwatch/domain/entities"
		"net/http"

		"github.com/gin-gonic/gin"
	)

	type CreateSmartwatchController struct {
		createSmartwatchUseCase *application.CreateSmartwatchUseCase
	}

	func NewCreateSmartwatchController(createSmartwatchUseCase *application.CreateSmartwatchUseCase) *CreateSmartwatchController {
		return &CreateSmartwatchController{
			createSmartwatchUseCase: createSmartwatchUseCase,
		}
	}

	func (c *CreateSmartwatchController) SaveSmartwatch(ctx *gin.Context) {
		var data entities.Smartwatch
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.createSmartwatchUseCase.Run(&data); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Smartwatch guardada",
			"data":    data,
		})
	}
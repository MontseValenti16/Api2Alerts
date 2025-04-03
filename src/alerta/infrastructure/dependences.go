package infrastructure

import (
	"LifeGuardAlertas/src/alerta/application"
	"LifeGuardAlertas/src/alerta/infrastructure/repository/mysql"
	"LifeGuardAlertas/src/alerta/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	db := mysql.NewMySQL()
	useCase := application.NewCreateAlertaUseCase(db)
	controller := handlers.NewCreateAlertaController(useCase)

	r.POST("/alertas", controller.SaveFiebre)
}

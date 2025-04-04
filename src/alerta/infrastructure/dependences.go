package infrastructure

import (
	"LifeGuardAlertas/src/alerta/application"
	"LifeGuardAlertas/src/alerta/infrastructure/repository/mysql"
	"LifeGuardAlertas/src/alerta/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	repo := mysql.NewMySQL()
	
	// Crear casos de uso
	createUC := application.NewCreateAlertaUseCase(repo)
	getUC := application.NewGetAlertasUseCase(repo)
	
	// Crear handlers
	createHandler := handlers.NewCreateAlertaController(createUC)
	getHandler := handlers.NewGetAlertaController(getUC)

	// Configurar rutas
	r.POST("/alertas", createHandler.SaveFiebre)
	r.GET("/alertas", getHandler.GetAllAlertas)
	r.GET("/alertas/:macaddress", getHandler.GetAlertasByMacAddress)
}

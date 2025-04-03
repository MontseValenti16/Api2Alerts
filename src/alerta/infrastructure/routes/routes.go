package infrastructure

import (
	"LifeGuardAlertas/src/alerta/application"
	"LifeGuardAlertas/src/alerta/infrastructure/handlers"
	"LifeGuardAlertas/src/alerta/infrastructure/repository/mysql"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Configurar repositorio
	db := mysql.NewMySQL()
	
	// Casos de uso
	createUseCase := application.NewCreateAlertaUseCase(db)
	getUseCase := application.NewGetAlertasUseCase(db)
	
	// Handlers
	createHandler := handlers.NewCreateAlertaController(createUseCase)
	getHandler := handlers.NewGetAlertaController(getUseCase)

	// Grupo de rutas para alertas
	alertGroup := r.Group("/alertas")
	{
		alertGroup.POST("/", createHandler.SaveFiebre)
		alertGroup.GET("/", getHandler.GetAllAlertas)
		alertGroup.GET("/:macaddress", getHandler.GetAlertasByMacAddress)
	}
}
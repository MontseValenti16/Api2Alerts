// src/samrtwatch/infrastrucutre/routes/routes.go
package routes

import (
	"LifeGuardAlertas/src/samrtwatch/application"
	"LifeGuardAlertas/src/samrtwatch/infrastrucutre/controller"
	"LifeGuardAlertas/src/samrtwatch/infrastrucutre/repository/mysql"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	repo := mysql.NewMySQL()

	createUC := application.NewCreateSmartwatchUseCase(repo)
	getUC := application.NewGetSmartwatchUseCase(repo)
	updateUC := application.NewUpdateSmartwatchUseCase(repo)

	controller := controller.NewSmartwatchController(createUC, getUC, updateUC)

	r.POST("/smartwatch", controller.CreateSmartwatch)
	r.GET("/smartwatch/:macaddress", controller.GetSmartwatch)
	r.PUT("/smartwatch/:macaddress", controller.UpdateSmartwatch)

	log.Println("Rutas inicializadas")
}
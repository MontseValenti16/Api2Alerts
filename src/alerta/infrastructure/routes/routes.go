package infraestructure

import (
	"LifeGuardAlertas/src/alerta/infraestructure/controller"
	"github.com/gin-gonic/gin"
)

func AlertaRoutes(r *gin.Engine, mpu6050Service *infraestructure.Mpu6050Controller) {
	mpu6050Group := r.Group("/alertas")
	{
		mpu6050Group.POST("/", mpu6050Service.CreateProduct)
	}
}

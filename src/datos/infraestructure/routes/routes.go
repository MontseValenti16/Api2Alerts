package routes

import (
	"LifeGuardAlertas/src/datos/application"
	handlers "LifeGuardAlertas/src/datos/infraestructure/handler"
	"LifeGuardAlertas/src/datos/infraestructure/repository/mysql"

	"github.com/gin-gonic/gin"
)

func InitDatosRoutes(r *gin.Engine) {
	// Repositorios
	repoAnuales := mysql.NewDatosAnualesMySQL()
	repoMensuales := mysql.NewDatosMensualesMySQL()
	repoSemanales := mysql.NewDatosSemanalesMySQL()
	repoDiarios := mysql.NewDatosDiariosMySQL()

	// Datos anuales
	anualesUC := application.NewGetAllDatosAnualesUseCase(repoAnuales)
	anualesHandler := handlers.NewGetAllDatosAnualesController(anualesUC)
	r.GET("/datos/anuales", anualesHandler.Handle)

	// Datos mensuales
	mensualesUC := application.NewGetAllDatosMensualesUseCase(repoMensuales)
	mensualesHandler := handlers.NewGetAllDatosMensualesController(mensualesUC)
	r.GET("/datos/mensuales", mensualesHandler.Handle)

	// Datos semanales
	semanalesUC := application.NewGetAllDatosSemanalUseCase(repoSemanales)
	semanalesHandler := handlers.NewGetAllDatosSemanalesController(semanalesUC)
	r.GET("/datos/semanales", semanalesHandler.Handle)

	// Datos diarios
	diariosUC := application.NewGetAllDatosDiariosUseCase(repoDiarios)
	diariosHandler := handlers.NewGetAllDatosDiariosController(diariosUC)
	r.GET("/datos/diarios", diariosHandler.Handle)
}

package main

import (
	"LifeGuardAlertas/src/alerta/infrastructure"
	"LifeGuardAlertas/src/core/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.MiddlewareCORS())
	infrastructure.InitRoutes(r)

	log.Println("Servidor iniciado en :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Error iniciando servidor: %v", err)
	}
}

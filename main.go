package main

import (
	"LifeGuardAlertas/src/core/middleware"
	"LifeGuardAlertas/src/alerta/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.MiddlewareCORS())
	infrastructure.InitRoutes(r)
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}

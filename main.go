package main

import (
    "LifeGuardAlertas/src/alerta/infrastructure"
    "LifeGuardAlertas/src/core/middleware"
    "LifeGuardAlertas/src/core/rabbitmq"  
    "log"

    infraSmart "LifeGuardAlertas/src/samrtwatch/infrastrucutre" 
    infraDatos "LifeGuardAlertas/src/datos/infraestructure/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializa la conexión MQTT
    if err := mqtt.InitMQTT(); err != nil { 
        log.Fatalf("Error al inicializar MQTT: %v", err)
    }

    r := gin.Default()
    r.Use(middleware.MiddlewareCORS())

    // Inicializa las rutas de alerta
    infrastructure.InitRoutes(r)

    // Inicializa las rutas de smartwatch
    infraSmart.InitRoutes(r)

    infraDatos.InitDatosRoutes(r)



    log.Println("Servidor iniciado en :8081")
    if err := r.Run(":8081"); err != nil {
        log.Fatalf("Error iniciando servidor: %v", err)
    }
}
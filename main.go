package main

import (
	"LifeGuardAlertas/src/core/middleware"
	"LifeGuardAlertas/src/core/rabbitmq"
	"LifeGuardAlertas/src/alerta/infrastructure"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configurar manejo de señales
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Inicializar MQTT
	if err := mqtt.InitMQTT(); err != nil {
		log.Fatalf("Error inicializando MQTT: %v", err)
	}
	defer mqtt.Close()

	// Crear router Gin
	r := gin.Default()
	r.Use(middleware.MiddlewareCORS())

	// Configurar rutas
	infrastructure.InitRoutes(r)

	// Iniciar servidor
	go func() {
		log.Println("Iniciando servidor en :8081")
		if err := r.Run(":8081"); err != nil {
			log.Fatalf("Error iniciando servidor: %v", err)
		}
	}()

	// Esperar señal de cierre
	<-sigs
	log.Println("Apagando servidor...")
}
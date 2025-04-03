package mqtt

import (
	"LifeGuardAlertas/src/alerta/domain/entities"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	AlertTopic = "cola_alertas" // Topic MQTT para las alertas
)

var (
	client mqtt.Client
)

// Configuración de MQTT (basada en tu ejemplo)
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

func getConfig() Config {
	return Config{
		Host:     "52.5.39.187", // Misma IP que en tu ejemplo
		Port:     "1883",        // Puerto estándar MQTT
		Username: "admin",       // Mismo usuario
		Password: "password",    // Misma contraseña
	}
}

func InitMQTT() error {
	config := getConfig()

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", config.Host, config.Port))
	opts.SetClientID("alertas-publisher")
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(5 * time.Second)
	opts.SetOnConnectHandler(onConnect)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("error conectando a MQTT: %v", token.Error())
	}

	log.Println("Conexión MQTT establecida para alertas")
	return nil
}

func onConnect(c mqtt.Client) {
	log.Println("Reconectado a MQTT")
}

func SendAlert(alerta *entities.Alerta) error {
	if !client.IsConnected() {
		return fmt.Errorf("cliente MQTT no conectado")
	}

	payload, err := json.Marshal(alerta)
	if err != nil {
		return fmt.Errorf("error serializando alerta: %v", err)
	}

	token := client.Publish(AlertTopic, 1, false, payload)
	token.Wait()
	if token.Error() != nil {
		return fmt.Errorf("error publicando alerta: %v", token.Error())
	}

	log.Printf("Alerta enviada por MQTT (Topic: %s): %+v", AlertTopic, alerta)
	return nil
}

func Close() {
	if client != nil && client.IsConnected() {
		client.Disconnect(250)
		log.Println("Conexión MQTT cerrada")
	}
}
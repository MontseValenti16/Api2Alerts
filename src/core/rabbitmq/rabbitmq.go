package mqtt

import (
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	RabbitMQHost     = "52.5.39.187" // O la IP del contenedor
	RabbitMQPort     = "1883"
	RabbitMQUser     = "guest"
	RabbitMQPassword = "guest"
	AlertTopic       = "test.#" // Topic MQTT
)

var client mqtt.Client

func InitMQTT() error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", RabbitMQHost, RabbitMQPort))
	opts.SetUsername(RabbitMQUser)
	opts.SetPassword(RabbitMQPassword)
	opts.SetClientID("api-alertas")
	opts.SetAutoReconnect(true)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("error conectando a RabbitMQ MQTT: %v", token.Error())
	}
	log.Println("✅ Conexión MQTT establecida con RabbitMQ")
	return nil
}

func SendAlert(alerta interface{}) error {
    payload, err := json.Marshal(alerta)
    if err != nil {
        return fmt.Errorf("error serializando alerta: %v", err)
    }

    // QoS 1 para entrega garantizada
    token := client.Publish(AlertTopic, 1, false, payload)
    token.Wait()
    
    if token.Error() != nil {
        return fmt.Errorf("error enviando alerta: %v", token.Error())
    }

    log.Printf("📩 Alerta enviada a RabbitMQ (Topic: %s, Payload: %s)", AlertTopic, string(payload))
    return nil
}
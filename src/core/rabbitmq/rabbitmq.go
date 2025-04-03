// core/rabbitmq/rabbitmq.go
package rabbitmq

import (
	"LifeGuardAlertas/src/alerta/domain/entities"
	"encoding/json"
	"fmt"
	"log"

)

const (
	QueueName = "cola_alertas"
)

var conn *amqp.Connection
var channel *amqp.Channel

func InitRabbitMQ() error {
	var err error
	conn, err = amqp.Dial("amqp://guest:guest@52.5.39.187:1883/")
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	_, err = channel.QueueDeclare(
		QueueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	return nil
}

func SendToRabbitMQ(alerta *entities.Alerta) error {
	if channel == nil {
		return fmt.Errorf("RabbitMQ channel not initialized")
	}

	body, err := json.Marshal(alerta)
	if err != nil {
		return fmt.Errorf("error marshaling alerta: %w", err)
	}

	err = channel.Publish(
		"",        // exchange
		QueueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	log.Printf("Alerta enviada a RabbitMQ: %+v", alerta)
	return nil
}

func Close() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}
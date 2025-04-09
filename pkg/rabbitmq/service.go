package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var (
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
)

func Init() {
	var err error
	Conn, err = amqp091.Dial("amqp://guest:guest@study-buddies-rabbitmq:5672/")
	if err != nil {
		log.Fatal("❌ Failed to connect to RabbitMQ:", err)
	}

	Channel, err = Conn.Channel()
	if err != nil {
		log.Fatal("❌ Failed to open channel:", err)
	}

	_, err = Channel.QueueDeclare(
		"class_notification", false, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("❌ Failed to declare queue:", err)
	}

	log.Println("✅ RabbitMQ connected and queue declared")
}

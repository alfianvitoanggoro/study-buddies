package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/robfig/cron"
)

func PublishMessage(queueName, message string) error {
	err := Channel.Publish(
		"", queueName, false, false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Println("❌ Failed to publish message:", err)
		return err
	}

	log.Println("📤 Published message:", message)
	return nil
}

func CronJob() {
	c := cron.New()
	c.AddFunc("@every 10m", func() {
		msg := "🔔 Check upcoming classes every 10m!"
		PublishMessage("class_notification", msg)
		log.Println("📤 CRON published reminder")
	})
	c.Start()
}

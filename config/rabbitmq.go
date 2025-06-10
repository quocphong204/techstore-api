package config

import (
	"fmt"
	"time"


	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQConn *amqp.Connection

func ConnectRabbitMQ() {
	var err error
	maxRetries := 10

	for i := 1; i <= maxRetries; i++ {
		RabbitMQConn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err == nil {
			fmt.Println("✅ Connected to RabbitMQ")
			return
		}

		fmt.Printf("⏳ Retry %d/%d - RabbitMQ not ready: %v\n", i, maxRetries, err)
		time.Sleep(3 * time.Second)
	}
}

package queue

import (
    "encoding/json"
    "log"
    amqp "github.com/rabbitmq/amqp091-go"
)

func SendToQueue(data map[string]interface{}) {
    conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
    if err != nil {
        log.Fatalf("❌ RabbitMQ connect failed: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("❌ RabbitMQ channel failed: %v", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare("order_queue", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("❌ Declare queue failed: %v", err)
    }

    body, _ := json.Marshal(data)
    err = ch.Publish("", q.Name, false, false, amqp.Publishing{
        ContentType: "application/json",
        Body:        body,
    })
    if err != nil {
        log.Fatalf("❌ Publish message failed: %v", err)
    }

    log.Println("📤 Gửi đơn hàng vào RabbitMQ thành công")
}

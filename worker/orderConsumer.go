package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"techstore-api/models"
	"techstore-api/utils"

	"github.com/rabbitmq/amqp091-go"
)

func StartOrderConsumer(ch *amqp091.Channel) {
	_, err := ch.QueueDeclare("order_queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal("❌ Failed to declare queue:", err)
	}

	msgs, err := ch.Consume("order_queue", "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("❌ Cannot consume:", err)
	}

	go func() {
		for msg := range msgs {
			var order models.Order
			if err := json.Unmarshal(msg.Body, &order); err != nil {
				log.Println("❌ Failed to parse order:", err)
				continue
			}

			fmt.Println("📦 Received order:", order.ID)

			subject := fmt.Sprintf("Xác nhận đơn hàng #%d", order.ID)
			body := fmt.Sprintf(`<h2>✅ Đơn hàng #%d đã được tiếp nhận!</h2><p>Tổng tiền: <b>%.2f VND</b></p>`, order.ID, order.Total)

			err := utils.SendOrderEmail("your_email@gmail.com", subject, body)
			if err != nil {
				log.Println("❌ Failed to send email:", err)
			} else {
				log.Println("📧 Email sent for order", order.ID)
			}
		}
	}()
}

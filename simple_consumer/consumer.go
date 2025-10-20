package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("failed to connect to rmq:", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("failed to open channel:", err)
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"q.hello", // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // dont wait for server response
		nil,
	)

	if err != nil {
		fmt.Println("error declairing queue:", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer name,
		true,   // auto acknowledgement
		false,  // exclusive
		false,  // no locale
		false,  // no wait
		nil,
	)
	

	if err != nil {
		fmt.Println("error consuming message:", err)
		return
	}

	for msg := range msgs {
		fmt.Println("got message:", string(msg.Body))
	}

}

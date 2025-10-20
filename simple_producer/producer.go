package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	//connect to rmq
	//                    user:password
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

	// declare queue
	q, err := ch.QueueDeclare(
		"q.hello", // queue name
		true, // durable
		false, // auto delete
		false, // exclusive
		false, // dont wait for server response
		nil,
	)

	if err != nil {
		fmt.Println("error declairing queue:", err)
		return
	}

	fmt.Println(q.Name)
	err = ch.Publish("",q.Name,false,false, amqp091.Publishing{
		ContentType: "text/plain",
		Body: []byte("whats up!, saying hi"),
	})

	if err != nil {
		fmt.Println("error publishing:", err)
		return
	}

}

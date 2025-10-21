package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	// if we set custom vertual host we write localhost:5672/{hostname}

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

	err = ch.ExchangeDeclare("x.topic", "topic", true, false, false, false, nil)

	if err != nil {
		fmt.Println("error creating exchange:", err)
		return
	}

	q, err := ch.QueueDeclare(
		"q.queue1", // queue name
		true,       // durable
		false,      // auto delete
		false,      // exclusive
		false,      // dont wait for server response
		amqp091.Table{
			"x-queue-type": "quorum",
		},
	)

	if err != nil {
		fmt.Println("error declairing queue:", err)
		return
	}

	q2, err := ch.QueueDeclare(
		"q.queue2", // queue name
		true,       // durable
		false,      // auto delete
		false,      // exclusive
		false,      // dont wait for server response
		amqp091.Table{
			"x-queue-type": "quorum",
		},
	)

	if err != nil {
		fmt.Println("error declairing queue:", err)
		return
	}

	// * only allows one word, # allows multiple or no word
	err = ch.QueueBind(q.Name, "*.error", "x.topic", false, nil)

	if err != nil {
		fmt.Println("error binding exchange:", err)
		return
	}

	err = ch.QueueBind(q2.Name, "kern.#", "x.topic", false, nil)

	if err != nil {
		fmt.Println("error binding exchange:", err)
		return
	}

	routingKeys := []string{"kern.log.smth", "log.smth.error", "log.error", "kern", "app.error", "app.log"}

	for _, key := range routingKeys {

		err = ch.Publish("x.topic", key, false, false, amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(key),
		})

		if err != nil {
			fmt.Println("error publishing:", err)
		}

	}
}

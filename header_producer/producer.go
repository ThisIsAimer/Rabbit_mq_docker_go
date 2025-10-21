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

	err = ch.ExchangeDeclare("x.header", "headers", true, false, false, false, nil)

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
	err = ch.QueueBind(q.Name, "", "x.header", false, amqp091.Table{
		"x-match": "all",
		"type":    "error",
		"format":  "json",
	})

	if err != nil {
		fmt.Println("error binding exchange:", err)
		return
	}

	err = ch.QueueBind(q2.Name, "", "x.header", false, amqp091.Table{
		"x-match": "any",
		"type":    "warning",
		"format":  "xml",
	})

	if err != nil {
		fmt.Println("error binding exchange:", err)
		return
	}

	messages := []struct {
		body    string
		headers amqp091.Table
	}{
		{
			body: "critical error in json",
			headers: amqp091.Table{
				"type":   "error",
				"format": "json",
			},
		}, {
			body: "warning alert in xml",
			headers: amqp091.Table{
				"type":   "warning",
				"format": "xml",
			},
		}, {
			body: "error alert in xml",
			headers: amqp091.Table{
				"type":   "error",
				"format": "xml",
			},
		},
		 {
			body: "warning in json",
			headers: amqp091.Table{
				"type":   "warning",
				"format": "json",
			},
		},
	}

	for _, message := range messages {

		err = ch.Publish("x.header", "", false, false, amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message.body),
			Headers: message.headers,
		})

		if err != nil {
			fmt.Println("error publishing:", err)
		}

	}
}

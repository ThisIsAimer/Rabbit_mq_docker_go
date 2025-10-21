package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	//connect to rmq
	//                    user:password
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

	err = ch.ExchangeDeclare("x.logs", "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Println("error creating exchange:", err)
		return
	}

	// declare queue
	q, err := ch.QueueDeclare(
		"q.hello", // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // dont wait for server response
		amqp091.Table{
			"x-queue-type": "quorum",
		},
	)

	if err != nil {
		fmt.Println("error declairing queue:", err)
		return
	}

	err = ch.QueueBind(q.Name, "test_key", "x.logs", false, nil)

	if err != nil {
		fmt.Println("error binding exchange:", err)
		return
	}

	// will stop program if message is not qued within 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("type a message to rmq (type quit to exit)")
	for {
		fmt.Print("enter a message:")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading string:", err)
			return
		}

		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "quit" {
			fmt.Println("exiting...")
			break
		}

		err = ch.PublishWithContext(ctx, "x.logs", "test_key", false, false, amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(input),
		})

		if err != nil {
			fmt.Println("error publishing:", err)
		}
	}

	_, err = ch.QueueDelete(
		q.Name, // queue name
		false,  // ifUnused (delete only if no consumers)
		false,  // ifEmpty (delete only if no messages)
		false,  // noWait
	)
	if err != nil {
		fmt.Println("Queue delete failed:", err)
		return
	}

	err = ch.ExchangeDelete(
		"x.logs", // exchange name
		false,    // ifUnused (delete only if no queues bound)
		false,    // noWait
	)

	if err != nil {
		fmt.Println("Queue delete failed:", err)
		return
	}

}

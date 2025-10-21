package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	var queueName string

	queueList := []string{"q.queue1", "q.queue2"}

	fmt.Println("queue options:", queueList)
	fmt.Print("enter queue name:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err != nil {
		fmt.Println("error reading string:", err)
		return
	}

	for _, name := range queueList {

		if name == input {
			queueName = name
		}

	}

	if queueName == "" {
		fmt.Println("invalid queue name")
		return
	}

	msgs, err := ch.Consume(
		queueName, // queue name
		"",        // consumer name,
		true,      // auto acknowledgement
		false,     // exclusive
		false,     // no locale
		false,     // no wait
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

package main

import (
	"bufio"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	failOnError(err,"Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	fmt.Println("enter message: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		body := scanner.Text()
		fmt.Println("sending: " + body)
		publishMessage(ch, body)
	}
}

func publishMessage(ch *amqp.Channel, body string) {
	err := ch.Publish(
		"logs_topic",
		severityForm(os.Args),
		false,
		false,
		amqp.Publishing{
			DeliveryMode:amqp.Persistent,
			ContentType:"text/plain",
			Body: []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func severityForm(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}

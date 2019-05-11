package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
    bytes := make([]byte, l)
    for i := 0; i < l; i++ {
    	bytes[i] = byte(randInt(65, 90))
	}

    return string(bytes)
}

func randInt(min, max int) int {
	return  min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)

	err = ch.Publish(
		"",
		"rpc_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			CorrelationId:corrId,
			ReplyTo:q.Name,
			Body:[] byte(strconv.Itoa(n)),
		})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer")
			break
		}
	}

	return
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("enter a number: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		n , err := strconv.Atoi(input)
		failOnError(err, "Failed to convert body to integer")
		log.Printf(" [x] Requesting fib(%d)", n)
		res, err := fibonacciRPC(n)

		log.Printf(" [.] Got %d", res)
	}

}

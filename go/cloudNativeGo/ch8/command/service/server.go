package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
)

func NewServer() *gin.Engine {
	deliveryDispatcher := initDispatcher("delivery")
	deliveryHandlers := NewHandlers(deliveryDispatcher)
	alertDispatcher := initDispatcher("alert")
	alertHandlers := NewHandlers(alertDispatcher)
	positionDispatcher := initDispatcher("position")
	positionHandlers := NewHandlers(positionDispatcher)
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	initRouter(engine, deliveryHandlers, alertHandlers, positionHandlers)

	return engine
}

func initDispatcher(name  string) queueDispatcher {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	failOnError(err,"Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()

	q, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")


	dispatcher := NewAmqpDispatcher(ch, q.Name, false)
	return dispatcher
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


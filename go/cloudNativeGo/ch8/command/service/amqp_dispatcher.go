package service

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type AmqpDispatcher struct {
	channel queuePublishableChanel
	queueName string
	mandatorySend bool
}

type queuePublishableChanel interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}

func NewAmqpDispatcher(publishChanel queuePublishableChanel, queueName string, mandatorySend bool) *AmqpDispatcher {
	return &AmqpDispatcher{publishChanel, queueName, mandatorySend}
}

func (q *AmqpDispatcher) DispatchMessage(message interface{}) (err error) {
	fmt.Printf("Dispatching message to queue %s\n", q.queueName)
	body, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("Failed to marshal message %v (%s) \n", message, err)
	}
	err = q.channel.Publish(
		"",
		q.queueName,
		q.mandatorySend,
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body:[]byte(body),
		})
	return
}

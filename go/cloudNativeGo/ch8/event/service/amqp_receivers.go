package service

import (
	"encoding/json"
	"fmt"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/common"
	"github.com/streadway/amqp"
	"log"
)

func DequeueEvent(alertChannel chan common.AlertSignalledEvent, telemetryChannel chan common.TelemetryUpdatedEvent, positionChanel chan common.PositionChangedEvent)  {
	fmt.Println("Starting AMQP queue de-serializer...")

	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	failOnError(err,"Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()

	alertsQ, err := ch.QueueDeclare(
		alertsQueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	positionQ, err := ch.QueueDeclare(
		positionsQueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	telemetryQ, err := ch.QueueDeclare(
		telemetryQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	alertsIn, _ := ch.Consume(
		alertsQ.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)

	positionIn, _ := ch.Consume(
		positionQ.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	telemetryIn, _ := ch.Consume(
		telemetryQ.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for {
			select {
			case alertRaw := <-alertsIn:
				dispatchAlert(alertRaw, alertChannel)
			case telemetryRaw := <-telemetryIn:
				dispatchTelemetry(telemetryRaw, telemetryChannel)
			case positionRaw := <-positionIn:
				dispatchPosition(positionRaw, positionChanel)
			}
		}
	}()

}

func dispatchAlert(alertRaw amqp.Delivery, out chan common.AlertSignalledEvent) {
	var event common.AlertSignalledEvent
	err := json.Unmarshal(alertRaw.Body, &event)
	failOnError(err, "Failed to de-serialize raw alert from queue")
	out <- event
	return
}

func dispatchTelemetry(telemetryRaw amqp.Delivery, out chan common.TelemetryUpdatedEvent) {
	var event common.TelemetryUpdatedEvent
	err := json.Unmarshal(telemetryRaw.Body, &event)
	failOnError(err, "Failed to de-serialize raw alert from queue")
	out <- event
	return
}

func dispatchPosition(positionRaw amqp.Delivery, out chan common.PositionChangedEvent) {
	var event common.PositionChangedEvent
	err := json.Unmarshal(positionRaw.Body, &event)
	failOnError(err, "Failed to de-serialize raw alert from queue")
	out <- event
	return
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}


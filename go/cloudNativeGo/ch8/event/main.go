package main

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/command/service"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/common"
	service2 "github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/event/service"
	"github.com/labstack/gommon/log"
)

func main() {
	repository, err := service.NewRepository()
	if err != nil {
		log.Fatalf("Failed to init mongodb repository: %s", err)
	}

	var alertChan chan common.AlertSignalledEvent
	var telemetryChan chan common.TelemetryUpdatedEvent
	var positionChan chan common.PositionChangedEvent

	service2.DequeueEvent(alertChan, telemetryChan, positionChan)

	service2.ConsumeEvents(alertChan, telemetryChan, positionChan, repository)
}

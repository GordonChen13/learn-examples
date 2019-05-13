package service

import (
	"fmt"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/common"
)

func ConsumeEvents(alertChanel chan common.AlertSignalledEvent, telemetryChanel chan common.TelemetryUpdatedEvent, positionChanel chan common.PositionChangedEvent, repo eventRepository) {
	go func() {
		fmt.Println("Started event consumer goroutine")
		for {
			select {
			case alert := <-alertChanel:
				processAlert(repo, alert)
			case telemetry := <-telemetryChanel:
				processTelemetry(repo, telemetry)
			case position := <-positionChanel:
				processPosition(repo, position)
			}
		}
	}()
}

func processAlert(repo eventRepository, alertEvent common.AlertSignalledEvent) {
	fmt.Printf("Processing alert %+v \n", alertEvent)

	repo.UpdateLastAlertEvent(alertEvent)
}

func processTelemetry(repo eventRepository, telemetryEvent common.TelemetryUpdatedEvent) {
	fmt.Printf("Processing alert %+v \n", telemetryEvent)

	repo.UpdateLastTelemetryEvent(telemetryEvent)
}

func processPosition(repo eventRepository, postitionEvent common.PositionChangedEvent) {
	fmt.Printf("Processing alert %+v \n", postitionEvent)

	repo.UpdateLastPositionEvent(postitionEvent)
}

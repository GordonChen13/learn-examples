package service

import "github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/common"

const (
	alertsQueueName = "alerts"
	telemetryQueueName = "telemetry"
	positionsQueueName = "positions"
)

type eventRepository interface {
	UpdateLastTelemetryEvent(telemetryEvent common.TelemetryUpdatedEvent) (err error)
	UpdateLastAlertEvent(alertEvent, event common.AlertSignalledEvent) (err error)
	UpdateLastPositionEvent(positionEvent common.PositionChangedEvent) (err error)
}

package common

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

type TelemetryUpdatedEvent struct {
	DroneID int64 `json:"drone_id"`
	RemainingBattery int `json:"battery"`
	Uptime int `json:"uptime"`
	CoreTemp int `json:"core_temp"`
	ReceiveOn int64 `json:"receive_on"`
}

type AlertSignalledEvent struct {
	DroneID int64 `json:"drone_id"`
	FaultCode int `json:"fault_code"`
	Description string `json:"description"`
	ReceiveOn int64 `json:"receive_on"`
}

type PositionChangedEvent struct {
	DroneID int64 `json:"drone_id"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude float32 `json:"altitude"`
	CurrentSpeed float32 `json:"current_speed"`
	HeadingCardinal int `json:"heading_cardinal"`
	ReceiveOn int64 `json:"receive_on"`
}

func NewDroneId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	id := node.Generate()
	return int64(id)
}

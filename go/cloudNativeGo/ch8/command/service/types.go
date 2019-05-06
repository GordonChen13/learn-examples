package service

type telemetryCommand struct {
	DroneID int64 `json:"drone_id"`
	RemainingBattery int `json:"battery"`
	Uptime int `json:"uptime"`
	CoreTemp int `json:"core_temp"`
}

type alertCommand struct {
	DroneID int64 `json:"drone_id"`
	FaultCode int `json:"fault_code"`
	Description string `json:"description"`
}

type positionCommand struct {
	DroneID int64 `json:"drone_id"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude float32 `json:"altitude"`
	CurrentSpeed float32 `json:"current_speed"`
	HeadingCardinal int `json:"heading_cardinal"`
}

type queueDispatcher interface {
	DispatchMessage(message interface{}) (err error)
}

func (t telemetryCommand) isValid() (valid bool) {
	valid = true
	if t.DroneID == 0 {
		valid = false
	}
	if t.Uptime == 0 {
		valid = false
	}
	return
}

func (a alertCommand) isValid() (valid bool) {
	valid = true
	if a.DroneID == 0 {
		valid = false
	}
	if len(a.Description) == 0 {
		valid = false
	}
	return
}

func (p positionCommand) isValid() (valid bool) {
	valid = true
	if p.DroneID == 0 {
		valid = false
	}
	if p.Longitude < 0 || p.Latitude < 0 || p.HeadingCardinal < 0 {
		valid = false
	}
	return
}

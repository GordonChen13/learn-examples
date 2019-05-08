package service

type telemetryCommand struct {
	DroneID int64 `json:"drone_id" form:"drone_id" binding:"required"`
	RemainingBattery int `json:"battery" form:"battery" binding:"required"`
	Uptime int `json:"uptime" form:"uptime" binding:"required"`
	CoreTemp int `json:"core_temp" form:"core_temp" binding:"required"`
}

type alertCommand struct {
	DroneID int64 `json:"drone_id" form:"drone_id" binding:"required"`
	FaultCode int `json:"fault_code" form:"fault_code" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type positionCommand struct {
	DroneID int64 `json:"drone_id" form:"drone_id" binding:"required"`
	Latitude float32 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" form:"longitude" binding:"required"`
	Altitude float32 `json:"altitude" form:"altitude" binding:"required"`
	CurrentSpeed float32 `json:"current_speed" form:"current_speed" binding:"required"`
	HeadingCardinal int `json:"heading_cardinal" form:"heading_cardinal" binding:"required"`
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

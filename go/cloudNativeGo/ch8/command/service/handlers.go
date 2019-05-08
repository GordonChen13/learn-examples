package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handlers struct {
	dispatch queueDispatcher
}

func NewHandlers(dispatch queueDispatcher) *Handlers{
	return &Handlers{dispatch}
}

func (h *Handlers) AddTelemetryCommand(ctx *gin.Context) {
	var form telemetryCommand
	if err := ctx.ShouldBind(&form); err != nil {
		println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.ParseInt(ctx.PostForm("drone_id"), 10, 64)
	battery, err := strconv.Atoi(ctx.PostForm("battery"))
	uptime, err := strconv.Atoi(ctx.PostForm("uptime"))
	temp, err := strconv.Atoi(ctx.PostForm("core_temp"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	command := &telemetryCommand{
		DroneID: id,
		RemainingBattery: battery,
		Uptime: uptime,
		CoreTemp: temp,
	}
	err = h.dispatch.DispatchMessage(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusCreated, command)
}

func (h *Handlers) AddAlertCommand(ctx *gin.Context)  {
	var form alertCommand
	if err := ctx.ShouldBind(&form); err != nil {
		println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.ParseInt(ctx.PostForm("drone_id"), 10, 64)
	faultCode, err := strconv.Atoi(ctx.PostForm("fault_code"))
	descriptin := ctx.PostForm("description")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	command := &alertCommand{
		DroneID: id,
		FaultCode:faultCode,
		Description:descriptin,
	}
	err = h.dispatch.DispatchMessage(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusCreated, command)
}

func (h *Handlers) AddPositionCommand(ctx *gin.Context) {
	var form positionCommand
	if err := ctx.ShouldBind(&form); err != nil {
		println(err.Error())
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.ParseInt(ctx.PostForm("drone_id"), 10, 64)
	altitude, err := strconv.ParseFloat(ctx.PostForm("altitude"), 32)
	longitude, err := strconv.ParseFloat(ctx.PostForm("longitude"), 32)
	latitude, err := strconv.ParseFloat(ctx.PostForm("latitude"), 32)
	speed, err := strconv.ParseFloat(ctx.PostForm("current_speed"), 32)
	cardinal, err := strconv.Atoi(ctx.PostForm("heading_cardinal"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	command := &positionCommand{
		DroneID: id,
		Altitude:float32(altitude),
		Latitude:float32(latitude),
		Longitude:float32(longitude),
		CurrentSpeed:float32(speed),
		HeadingCardinal:cardinal,
	}
	err = h.dispatch.DispatchMessage(command)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusCreated, command)
}
package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handlers struct {
	dispatch queueDispatcher
}

func NewHandlers(dispatch queueDispatcher) *Handlers{
	return &Handlers{dispatch}
}

func (h *Handlers) AddTelemetryCommand(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, "")
}

func (h *Handlers) AddAlertCommand(ctx *gin.Context)  {

}

func (h *Handlers) AddPositionCommand(ctx *gin.Context) {

}
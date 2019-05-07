package service

import "github.com/gin-gonic/gin"

func initRouter(router *gin.Engine, handler *Handlers)  {
	router.POST("/api/command/telemetry", handler.AddTelemetryCommand)
	router.POST("/api/command/alerts", handler.AddAlertCommand)
	router.POST("/api/command/positions", handler.AddPositionCommand)
}
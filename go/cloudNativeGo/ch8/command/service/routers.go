package service

import "github.com/gin-gonic/gin"

func initRouter(router *gin.Engine, delivery, alert, position *Handlers)  {
	router.POST("/api/command/telemetry", delivery.AddTelemetryCommand)
	router.POST("/api/command/alerts", alert.AddAlertCommand)
	router.POST("/api/command/positions",position.AddPositionCommand)
}
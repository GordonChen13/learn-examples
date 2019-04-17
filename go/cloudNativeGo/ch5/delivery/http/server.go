package http

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	
	router.GET("/test", TestHandler)
	router.POST("/match", CreateMatch)
	
	return router
}



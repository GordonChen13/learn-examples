package service

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	router := gin.Default()
	
	router.GET("/test", TestHandler)
	
	return router
}



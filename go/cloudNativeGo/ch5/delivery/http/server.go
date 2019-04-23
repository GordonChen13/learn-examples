package http

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	return router
}

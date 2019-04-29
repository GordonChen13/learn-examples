package service

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	server := gin.Default()
	NewHandler(server)
	return server
}
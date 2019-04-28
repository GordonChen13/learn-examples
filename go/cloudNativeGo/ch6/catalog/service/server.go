package service

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	server := gin.Default()
	return server
}
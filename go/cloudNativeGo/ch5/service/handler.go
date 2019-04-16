package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is a test")
}

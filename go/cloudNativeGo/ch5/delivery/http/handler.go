package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is a test")
}

func CreateMatch(c *gin.Context) {
	name := c.PostForm("name")
	c.String(http.StatusCreated, name)
}

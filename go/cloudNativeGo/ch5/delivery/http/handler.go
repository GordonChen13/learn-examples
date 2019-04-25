package main

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MatchHandler struct {
	useCase usecase.Match
}

func NewMatchHandler(router *gin.Engine, matchUseCase usecase.Match) {
	handler := &MatchHandler{
		useCase: matchUseCase,
	}

	router.GET("/test", handler.Test)
	router.POST("/match", handler.CreateMatch)
	router.GET("/match", handler.QueryMatch)
}

func (m *MatchHandler) Test(c *gin.Context) {
	c.String(http.StatusOK, "This is a test")
}

func (m *MatchHandler) CreateMatch(c *gin.Context) {
	name := c.PostForm("name")

	res, err := m.useCase.Create(c, name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, res)
}

func (m *MatchHandler) QueryMatch(c *gin.Context)  {
	name := c.Param("name")
	if name != "" {
		m.GetMatchByName(c)
	}

	id := c.Param("id")
	if id!= "" {
		m.GetMatchById(c)
	}
}

func (m *MatchHandler) GetMatchByName(c *gin.Context) {
	name := c.Param("name")

	res, err := m.useCase.GetByName(c, name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, res)
}

func (m *MatchHandler) GetMatchById(c *gin.Context)  {
	id := c.Param("id")

	res, err := m.useCase.GetById(c, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, res)
}

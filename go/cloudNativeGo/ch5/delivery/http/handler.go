package http

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/models"
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
}

func (m *MatchHandler) Test(c *gin.Context) {
	c.String(http.StatusOK, "This is a test")
}

func (m *MatchHandler) CreateMatch(c *gin.Context) {
	name := c.PostForm("name")
	id := "232324"

	match := &models.Match{
		Id:    id,
		Name:  name,
		Moves: nil,
	}
	res, err := m.useCase.Create(c, match)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, res)
}

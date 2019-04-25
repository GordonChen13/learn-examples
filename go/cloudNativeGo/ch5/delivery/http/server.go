package main

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/repository/impl/mongo"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/usecase/impl"
	"github.com/gin-gonic/gin"
	"log"
)

func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	matchRepository, err := mongo.NewMatchRepository()
	if err != nil {
		log.Fatalf("connect to mongodb fail, %s", err)
	}

	matchUseCase := impl.NewMatchUseCase(matchRepository)
	NewMatchHandler(router, matchUseCase)

	return router
}

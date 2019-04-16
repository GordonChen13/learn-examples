package main

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/service"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	router := service.NewServer()
	router .Run(":" + port)
}

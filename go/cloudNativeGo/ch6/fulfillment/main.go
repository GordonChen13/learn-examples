package main

import "github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch6/catalog/service"

const port string = "3000"
func main() {
	server := service.NewServer()

	server.Run(":" + port)
}

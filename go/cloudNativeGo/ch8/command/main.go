package main

import "github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch8/command/service"

func main()  {
	server := service.NewServer()

	server.Run(":8080")
}

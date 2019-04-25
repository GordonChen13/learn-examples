package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	port := viper.GetString(`server.address`)

	if len(port) == 0 {
		port = "3000"
	}

	router := NewServer()
	err := router.Run(":" + port)

	if err != nil {
		log.Fatalf("gin run error, %s", err)
	}
}

func init() {
	viper.SetConfigFile(`../../config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}


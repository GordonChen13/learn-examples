package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	address := viper.GetString(`server.address`)

	if len(address) == 0 {
		address = "3000"
	}

	router := NewServer()
	err := router.Run(address)

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


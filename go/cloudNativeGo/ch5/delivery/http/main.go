package http

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	port := viper.GetString(`server.address`)

	if len(port) == 0 {
		port = "3000"
	}

	router := NewServer()
	router.Run(":" + port)
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


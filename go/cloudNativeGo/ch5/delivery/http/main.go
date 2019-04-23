package http

import (
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	router := NewServer()
	router.Run(":" + port)
}

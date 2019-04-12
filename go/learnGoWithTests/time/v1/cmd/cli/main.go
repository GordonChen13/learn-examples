package main

import (
	"fmt"
	"github.com/GordonChen13/learn-examples/go/learnGoWithTests/time/v1"
	"github.com/labstack/gommon/log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewCLI(store, os.Stdin, poker.BlindAlerterFunc(poker.StdOutAlerter))
	game.PlayPoker()
}

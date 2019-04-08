package main

import (
	"github.com/GordonChen13/learn-examples/go/learnGoWithTests/commandLine/v1"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}

	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
	}
}

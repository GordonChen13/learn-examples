package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
	score map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.score[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.score[name]++
	return
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{map[string]int{}}}
	if err := http.ListenAndServe(":5000", server); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer)GetPlayerScore(name string) int {
	return p.store.GetPlayerScore(name)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request)  {
	player := r.URL.Path[len("/players/"):]

	score := p.GetPlayerScore(player)

	if score == 0{
		w.WriteHeader(http.StatusNotFound)
	}


	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request)  {
	player := r.URL.Path[len("/players/"):]
	p.store.RecordWin(player)

	w.WriteHeader(http.StatusAccepted)
}
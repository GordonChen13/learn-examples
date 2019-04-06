package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerServer struct {
	store PlayerStore
    http.Handler
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

type Player struct {
	Name string
	Wins int
}

const jsonContentType = "application/json"

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}


func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) getLeagueTable() []Player {
	return []Player{
		{"Chris", 20},
	}

}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
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
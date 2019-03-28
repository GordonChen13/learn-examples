package main

import (
	"fmt"
	"net/http"
)

type PlayerServer struct {
	store PlayerStore
	router *http.ServeMux
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store,
		http.NewServeMux(),
	}

	p.router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	return p
}


func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
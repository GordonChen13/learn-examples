package main

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type StubPlayerStore struct {
	scores map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int{
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

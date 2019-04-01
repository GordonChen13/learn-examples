package main
type InMemoryPlayerStore struct {
	score map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.score[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.score[name]++
	return
}

func (i *InMemoryPlayerStore) GetLeague() []Player{
	var league []Player
	for name, wins := range i.score {
		league = append(league, Player{name, wins})
	}
	return league
}

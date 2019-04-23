package models

type Match struct {
	Id string
	Name string
	Moves []Step
}

type Step struct {
	X int
	Y int
}

func NewMatchId() string {
	return "123456"
}

package models

type Match struct {
	id string
	name string
	moves []Step
}

type Step struct {
	x int
	y int
}

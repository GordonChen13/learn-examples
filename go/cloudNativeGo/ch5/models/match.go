package models

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

type Match struct {
	Id int64
	Name string
	Moves []Step
}

type Step struct {
	X int
	Y int
}

func NewMatchId() int64 {
    node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

    id := node.Generate()
    return int64(id)
}

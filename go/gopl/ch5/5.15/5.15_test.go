package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	cs := []struct{
		name string
		input []int
		want int
	} {
		{"return 0 when empty", []int{}, 0},
		{"one int", []int{-1}, -1},
		{"multi int", []int{-100, 1, 10}, 10},
		{"multi int2", []int{-100, -1, -10}, -1},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got := Max(c.input...)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestMin(t *testing.T) {
	cs := []struct{
		name string
		input []int
		want int
	} {
		{"return 0 when empty", []int{}, 0},
		{"one int", []int{-1}, -1},
		{"multi int", []int{-100, 1, 10}, -100},
		{"multi int2", []int{100, 1, 10}, 1},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got := Min(c.input...)
			assert.Equal(t, c.want, got)
		})
	}
}

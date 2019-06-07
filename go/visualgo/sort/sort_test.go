package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubble(t *testing.T) {
	cases := [] struct{
		name string
		input []int
		want []int
	}{
		{"small int", []int{1, 8, 10, 7, 3, 88, 56, 93}, []int{1, 3, 7, 8, 10, 56, 88, 93}},
		{"big int", []int{10000, 8, 10, 7000, 3, 88, 5600000, 93}, []int{3, 8, 10, 88, 93, 7000, 10000, 5600000}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := Bubble(c.input)
			assert.Equal(t, c.want, get)
		})
	}
}

func TestSelect(t *testing.T) {
	cases := [] struct{
		name string
		input []int
		want []int
	}{
		{"small int", []int{1, 8, 10, 7, 3, 88, 56, 93}, []int{1, 3, 7, 8, 10, 56, 88, 93}},
		{"big int", []int{10000, 8, 10, 7000, 3, 88, 5600000, 93}, []int{3, 8, 10, 88, 93, 7000, 10000, 5600000}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := Select(c.input)
			assert.Equal(t, c.want, get)
		})
	}
}

func TestInsert(t *testing.T) {
	cases := [] struct{
		name string
		input []int
		want []int
	}{
		{"small int", []int{1, 8, 10, 7, 3, 88, 56, 93}, []int{1, 3, 7, 8, 10, 56, 88, 93}},
		{"big int", []int{10000, 8, 10, 7000, 3, 88, 5600000, 93}, []int{3, 8, 10, 88, 93, 7000, 10000, 5600000}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := Insert(c.input)
			assert.Equal(t, c.want, get)
		})
	}
}

func TestMerge(t *testing.T) {
	cases := [] struct{
		name string
		input []int
		want []int
	}{
		{"small int", []int{1, 8, 10, 7, 3, 88, 56, 93}, []int{1, 3, 7, 8, 10, 56, 88, 93}},
		{"big int", []int{10000, 8, 10, 7000, 3, 88, 5600000, 93}, []int{3, 8, 10, 88, 93, 7000, 10000, 5600000}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			get := MergeSort(c.input)
			assert.Equal(t, c.want, get)
		})
	}
}

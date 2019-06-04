package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasSameChar(t *testing.T) {
	cs := []struct {
		name string
		input1 string
		input2 string
		want bool
	}{
		{"return true when input 'abcdef123' and '321fedbca'", "abcdef123", "321fedbca", true},
		{"return true when input 'abcccdef11123' and '321cfec1db1ca'", "abcccdef11123", "321cfec1db1ca", true},
		{"return false when input 'abbccdef11123' and '321cfec1db1ca'", "abbccdef11123", "321cfec1db1ca", false},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			get := HasSameChar(c.input1, c.input2)
			assert.Equal(t, c.want, get)
		})
	}
}

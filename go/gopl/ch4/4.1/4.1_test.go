package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiffBitCount(t *testing.T) {
	var b1, b2 = byte('a'), byte('b')
	fmt.Printf("%b, %b", b1,b2)
	got := DiffBitCount(b1, b2)
	assert.Equal(t, 2, got)
}

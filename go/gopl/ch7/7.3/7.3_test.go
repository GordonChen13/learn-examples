package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_String(t *testing.T) {
	tree := &tree{value: 6}
	add(tree, 1)
	add(tree, 5)
	add(tree, 2)
	add(tree, 4)
	add(tree, 3)

	str := tree.String()
	assert.Equal(t, "1 2 3 4 5 6", str)
}

package intset_test

import (
	"github.com/GordonChen13/learn-examples/go/gopl/ch6/intset"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSet_Elems(t *testing.T) {
	set := &intset.IntSet{}
	set.AddAll(1,2,3,4,5)
	got := set.Elems()

	assert.Equal(t, []int{1,2,3,4,5}, got)
}

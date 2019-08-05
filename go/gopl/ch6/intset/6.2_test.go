package intset_test

import (
	"github.com/GordonChen13/learn-examples/go/gopl/ch6/intset"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSet_AddAll(t *testing.T) {
	set3 := &intset.IntSet{}
	set3.AddAll(1, 10, 100)

	assert.Equal(t,true, set3.Has(1))
	assert.Equal(t,true, set3.Has(10))
	assert.Equal(t,true, set3.Has(100))
}

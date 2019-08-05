package intset_test

import (
	"github.com/GordonChen13/learn-examples/go/gopl/ch6/intset"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSet_IntersectWith(t *testing.T) {
	set1 := &intset.IntSet{}
	set1.AddAll(1,5,6,7,9)

	set2 := &intset.IntSet{}
	set2.AddAll(2,4,5,6,8,10,12)

	set1.IntersectWith(set2)
	assert.Equal(t, false, set1.Has(1), "should't have 1")
	assert.Equal(t, false, set1.Has(2), "should't have 2")
	assert.Equal(t, true, set1.Has(5), "should have 5")
	assert.Equal(t, true, set1.Has(6), "should have 6")
	assert.Equal(t, false, set1.Has(7), "should't have 7")
	assert.Equal(t, false, set1.Has(8), "should't have 8")
}

func TestIntSet_DifferenceWith(t *testing.T) {
	set1 := &intset.IntSet{}
	set1.AddAll(1,5,6,7,9)

	set2 := &intset.IntSet{}
	set2.AddAll(2,4,5,6,8,10,12)

	set1.DifferenceWith(set2)
	assert.Equal(t,true, set1.Has(1), "should have 1")
	assert.Equal(t, false, set1.Has(2), "should't have 2")
	assert.Equal(t,false, set1.Has(5), "should't have 5")
	assert.Equal(t,false, set1.Has(6), "should't have 6")
	assert.Equal(t,true, set1.Has(7), "should have 7")
	assert.Equal(t, false, set1.Has(8), "should't have 8")
	assert.Equal(t,true, set1.Has(9), "should have 9")
}

func TestIntSet_SymmetricWith(t *testing.T) {
	set1 := &intset.IntSet{}
	set1.AddAll(1,5,6,7,9)

	set2 := &intset.IntSet{}
	set2.AddAll(2,4,5,6,8,10,12)

	set1.SymmetricWith(set2)
	assert.Equal(t,true, set1.Has(1), "should have 1")
	assert.Equal(t,true, set1.Has(2), "should have 2")
	assert.Equal(t,false, set1.Has(5), "should't have 5")
	assert.Equal(t,false, set1.Has(6), "should't have 6")
	assert.Equal(t,true, set1.Has(7), "should have 7")
	assert.Equal(t,true, set1.Has(8), "should have 8")
	assert.Equal(t,true, set1.Has(10), "should have 10")
}

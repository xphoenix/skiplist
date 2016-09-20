package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xphoenix/linkedlist"
	"github.com/xphoenix/skiplist"
)

func idxLevel(sk *skiplist.SkipList, level int) *skiplist.IndexNode {
	p := sk.IndexHead
	for l := sk.Height - 1; l > level && p.Down != nil; p, l = p.Down, l-1 {

	}
	return p
}

///////////////////////////////////////////////////////////////////////////////
// SearchIndex tests
///////////////////////////////////////////////////////////////////////////////
func TestSearchIndexInHeadColumn(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(10, 8)
	dump("TestSearchInHeadColumn.dot", sk)

	n1, n2 := sk.SearchIndex(0, 0)
	assert.Equal(idxLevel(sk, 1), n1, "Left")
	assert.Equal(20, n2.Root.Value(), "Right")
}

func TestSearchIndexInHeadColumnToUpperLevel(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(10, 8)
	dump("TestSearchInHeadColumnToUpperLevel.dot", sk)

	n1, n2 := sk.SearchIndex(0, 2)
	assert.Equal(idxLevel(sk, 2), n1, "Left")
	assert.Equal(40, n2.Root.Value(), "Right")
}

func TestSearchIndexEndsInTailColumn(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(10, 8)
	dump("TestSearchEndsInTailColumn.dot", sk)

	n1, n2 := sk.SearchIndex(80, 1)
	assert.Equal(80, n2.Root.Value(), "Right")
	assert.Equal(n2, linkedlist.LoadState(n1).Next, "Left")
}

///////////////////////////////////////////////////////////////////////////////
// Search tests
///////////////////////////////////////////////////////////////////////////////
func TestSearchInHeadColumn(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(10, 8)
	dump("TestSearchInHeadColumn.dot", sk)

	n1, n2 := sk.Search(-1)
	assert.Equal(sk.DataHead, n1, "Left")
	assert.Equal(0, n2.Value(), "Right")
}

func TestSearchInTailColumn(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(10, 8)
	dump("TestSearchInHeadColumn.dot", sk)

	n1, n2 := sk.Search(95)
	assert.Equal(90, n1.Value(), "Left")
	assert.Equal(sk.DataTail, n2, "Right")
}

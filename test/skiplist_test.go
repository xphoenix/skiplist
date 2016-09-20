package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xphoenix/linkedlist"
	"github.com/xphoenix/skiplist"
)

///////////////////////////////////////////////////////////////////////////////
// New tests
///////////////////////////////////////////////////////////////////////////////
func TestNewSkipList(t *testing.T) {
	assert := assert.New(t)
	sk := skiplist.New(IntNodeComparator, 8)
	dump("TestNewSkipList.before.dot", sk)

	// Check index column
	h := 1
	for p := sk.IndexHead; p.Down != nil; p, h = p.Down, h+1 {
		assert.Equal(sk.DataHead, p.Root, "Index heads point to the data head")
		assert.Equal(sk.IndexTail, linkedlist.LoadState(p).Next, "Index heads linked to the shared tail")
	}
	assert.Equal(7, h, "Index column height")

	assert.Equal(sk.DataTail, sk.IndexTail.Root, "Index tail ppoints to the data tail")
	assert.Equal(sk.DataTail, linkedlist.LoadState(sk.DataHead).Next, "Data heads linked to the data tail")
}

///////////////////////////////////////////////////////////////////////////////
// Comparator tests
///////////////////////////////////////////////////////////////////////////////
func TestHeadsAreAlwaysSmallest(t *testing.T) {
	assert := assert.New(t)
	sk := skiplist.New(IntNodeComparator, 8)

	for p := sk.IndexHead; p.Down != nil; p = p.Down {
		assert.Equal(1, sk.Compare(math.MinInt64, p), "All index head are smallest values")
	}
	assert.Equal(1, sk.Compare(math.MinInt64, sk.DataHead), "Data head is the smallest value")
}

func TestTailsAreAlwaysGreater(t *testing.T) {
	assert := assert.New(t)
	sk := skiplist.New(IntNodeComparator, 8)

	assert.Equal(-1, sk.Compare(math.MaxInt64, sk.IndexTail), "Index tail is the gretest value")
	assert.Equal(-1, sk.Compare(math.MaxInt64, sk.DataTail), "Data tail is the greatest value")
}

func TestNodesCompareByUserComparator(t *testing.T) {
	assert := assert.New(t)
	sk := skiplist.New(IntNodeComparator, 8)

	assert.Equal(-1, sk.Compare(10, NewIntNode(20)))
	assert.Equal(1, sk.Compare(30, NewIntNode(0)))
	assert.Equal(0, sk.Compare(0, NewIntNode(0)))
}

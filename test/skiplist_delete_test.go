package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xphoenix/linkedlist"
	"github.com/xphoenix/skiplist"
)

///////////////////////////////////////////////////////////////////////////////
// Insert tests
///////////////////////////////////////////////////////////////////////////////
func TestDeleteSimple(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(8, 3)
	dump("TestDeleteSimple.before.dot", sk)

	n, suc, self := sk.Delete(10)
	assert.True(suc, "Deleted")
	assert.True(self, "By current thread")
	assert.Equal(n.Value(), 10, "Current node")
	assert.Equal(20, linkedlist.LoadState(n).Next.(skiplist.Node).Value(), "Current node has next link")

	n1, n2 := sk.Search(10)
	assert.Equal(0, n1.Value(), "Predecessor")
	assert.Equal(20, n2.Value(), "Successor")
	dump("TestDeleteSimple.after.dot", sk)
}

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
func TestInsertIntoEmpty(t *testing.T) {
	assert := assert.New(t)
	sk := skiplist.New(IntNodeComparator, 8)
	dump("TestInsertIntoEmpty.before.dot", sk)

	new := NewIntNode(10)
	n, status := sk.Insert(new)
	assert.True(status, "Insert success")
	assert.Equal(new, n, "Inserted node returns")
	dump("TestInsertIntoEmpty.after.dot", sk)
}

func TestInsertSimple(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(10, 4)
	dump("TestInsertSimple.before.dot", sk)

	new := NewIntNode(11)
	n, status := sk.Insert(new)
	assert.True(status, "Insert success")
	assert.Equal(20, linkedlist.LoadState(n).Next.(skiplist.Node).Value(), "Successor")
	dump("TestInsertSimple.after.dot", sk)
}

func TestInsertExisting(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(10, 4)
	dump("TestInsertExisting.before.dot", sk)

	new := NewIntNode(10)
	n, status := sk.Insert(new)
	assert.False(status, "Insert fails")
	assert.NotEqual(n, new, "Existing node returned")
	assert.Equal(linkedlist.LoadState(linkedlist.LoadState(sk.DataHead).Next).Next, n, "Node")
	dump("TestInsertExisting.after.dot", sk)
}

func TestInsertIntoFreezed(t *testing.T) {
	assert := assert.New(t)
	sk := NewIntSkipList(10, 6)

	_, b := sk.Search(60)
	state := linkedlist.LoadState(b)
	state.Flags = linkedlist.FREEZE
	dump("TestInsertIntoDeleted.before.dot", sk)

	new := NewIntNode(65)
	n, status := sk.Insert(new)
	assert.True(status, "Insert success")
	assert.Equal(new, n, "Node returned")
	assert.Equal(linkedlist.LoadState(b).Next, new, "Predecessor")
	assert.Equal(linkedlist.LoadState(n).Next, linkedlist.LoadState(state.Next).Next, "Successor")
	assert.Equal(linkedlist.DELETE, linkedlist.LoadState(state.Next).Flags, "Removed flag")
	assert.Equal(b, linkedlist.LoadState(state.Next).Back, "Removed backlink")
	dump("TestInsertIntoDeleted.after.dot", sk)
}

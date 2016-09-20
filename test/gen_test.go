package test

import (
	"github.com/xphoenix/linkedlist"
	"github.com/xphoenix/skiplist"
)

// NewIntSkipList generate ideally balanced skiplist with given height & and size
func NewIntSkipList(size, height int) *skiplist.SkipList {
	column, sk := make([]skiplist.Node, height, height), skiplist.New(IntNodeComparator, height)
	column[height-1] = sk.IndexHead
	column[0] = sk.DataHead
	for i := height - 1; i > 1; i-- {
		column[i-1] = column[i].(*skiplist.IndexNode).Down
	}

	for i := 0; i < size; i++ {
		node, level := NewIntNode(i*10), 0
		for j := i; j > 0 && j&1 == 0 && level < height; j = j >> 1 {
			level++
		}

		// Insert data
		linkedlist.Insert(column[0], node)
		column[0] = node

		// Insert index
		for j, down := 1, (*skiplist.IndexNode)(nil); j <= level; j++ {
			idx := skiplist.NewIndexNode(node, down)
			linkedlist.Insert(column[j], idx)
			column[j] = idx
			down = idx
		}
	}
	return sk
}

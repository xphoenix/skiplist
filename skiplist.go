package skiplist

import (
	"math/rand"
	"sync/atomic"

	"github.com/xphoenix/linkedlist"
)

// Node represents a single object inserted into the skiplist
// TODO: Value() can't be nil
type Node interface {
	linkedlist.Node
	Value() interface{}
}

// Comparator is a function used to compare two user supplied values. It returns:
// - -1 if left is less then right
// -  0 in case of equality
// - +1 whenever right happens to be greater then left
type Comparator func(left, right interface{}) int

// SkipList represents a single skiplist data structure instance. It is threadsafe
// so multiple go routines could write/read from it concurrently
type SkipList struct {
	cmp  Comparator
	seed int32

	Height    int
	IndexHead *IndexNode
	IndexTail *IndexNode
	DataHead  Node
	DataTail  Node
}

// New creates new empty skiplist with given maximum index height
func New(cmp Comparator, height int) *SkipList {
	// Create empty skiplist
	dTail := &StabNode{state: &linkedlist.State{}}
	iTail := &IndexNode{state: &linkedlist.State{}, Root: dTail, Down: nil}
	result := &SkipList{
		cmp:       cmp,
		seed:      rand.Int31(),
		Height:    height,
		IndexHead: nil,
		IndexTail: iTail,
		DataHead:  &StabNode{state: &linkedlist.State{Next: dTail}},
		DataTail:  dTail,
	}

	// Initialize head index column
	for down, i := (*IndexNode)(nil), 1; i < height; i++ {
		result.IndexHead = &IndexNode{
			state: &linkedlist.State{Next: iTail},
			Root:  result.DataHead,
			Down:  down,
		}
		down = result.IndexHead
	}
	return result
}

// Compare given value and skiplist node. Function uses user supplied compare
// function and ensures that technical head/tail nodes are always smaller/greater
// then any possible value
//
// Semantics of the function is the same as for user supplied comparator
func (s *SkipList) Compare(value interface{}, node Node) int {
	// TODO: something more performance? one more "if" freaks me out
	idx, ok := node.(*IndexNode)
	if ok {
		node = idx.Root
	}

	switch node {
	// TODO: replace special tail nodes by nil

	case s.IndexHead: // head node
		return 1
	case s.IndexTail: // tail "node"
		return -1

	case s.DataHead: // head node
		return 1
	case s.DataTail: // tail "node"
		return -1

	default: // intermediate nodes
		return s.cmp(value, node.Value())
	}
}

// RandomLevel generate pseudo random number in range [0, Height] where Height
// is the maximum skip list index height
func (s *SkipList) RandomLevel() int {
	mask := 0
	for mask == 0 {
		src := atomic.LoadInt32(&s.seed)
		x := src
		x ^= x << 13
		x ^= int32(uint32(x) >> 17)
		x ^= x << 5
		if atomic.CompareAndSwapInt32(&s.seed, src, x) {
			mask = int(x)
		}
	}

	if mask&0x80000001 == 0x80000001 {
		return 0
	}

	level := 1
	for ; mask&1 == 1 && level < s.Height; mask = mask >> 1 {
		level++
	}
	return level
}

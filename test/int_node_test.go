package test

import (
	"fmt"

	"github.com/xphoenix/linkedlist"
)

// IntNode is a basic implementatio of the linkedlist node
type IntNode struct {
	state *linkedlist.State
	value int
}

// IntNodeComparator comparator for skiplist
var IntNodeComparator = func(vl, vr interface{}) int {
	left, right := vl.(int), vr.(int)
	if left < right {
		return -1
	} else if left > right {
		return 1
	} else {
		return 0
	}
}

// NewIntNode create new int linked list node
func NewIntNode(value int) *IntNode {
	return &IntNode{
		state: &linkedlist.State{Next: nil, Back: nil, Flags: linkedlist.NONE},
		value: value,
	}
}

// State implements skiplist.Node interface
func (n *IntNode) State() **linkedlist.State {
	return &n.state
}

// Value implements skiplist.Node interface
func (n *IntNode) Value() interface{} {
	return n.value
}

// String implements Stringer interface
func (n *IntNode) String() string {
	return fmt.Sprintf("%d | %s", n.value, n.state.Flags)
}

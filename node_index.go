package skiplist

import (
	"fmt"

	"github.com/xphoenix/linkedlist"
)

// IndexNode represents a single node in the skiplist index towers. Node carries
// links to next node on the same index lever, link to bottom level node and a
// fast link to the data level
type IndexNode struct {
	state *linkedlist.State
	Down  *IndexNode
	Root  Node
}

// NewIndexNode creates a new index node
func NewIndexNode(root Node, down *IndexNode) *IndexNode {
	return &IndexNode{
		state: &linkedlist.State{Next: nil, Back: nil, Flags: linkedlist.NONE},
		Down:  down,
		Root:  root,
	}
}

// State implements skiplist Node
func (n *IndexNode) State() **linkedlist.State {
	return &n.state
}

// Value implements skiplist Node
func (n *IndexNode) Value() interface{} {
	return n.Root.Value()
}

// String implements Stringer interface
func (n *IndexNode) String() string {
	return fmt.Sprintf("%d", n.state.Flags)
}

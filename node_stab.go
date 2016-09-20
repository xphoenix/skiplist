package skiplist

import "github.com/xphoenix/linkedlist"

// StabNode represents an empty skipklist node that carry no useful content. Used
// as a marker  nodes for head/tail
type StabNode struct {
	state *linkedlist.State
}

// State implements skiplist Node
func (h *StabNode) State() **linkedlist.State {
	return &h.state
}

// Value implements skiplist Node
func (h *StabNode) Value() interface{} {
	return nil
}

// String implements Stringer interface
func (h *StabNode) String() string {
	return "stab"
}

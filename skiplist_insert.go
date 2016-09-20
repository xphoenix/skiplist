package skiplist

import "github.com/xphoenix/linkedlist"

// Insert new node into the skiplist. Note that all nodes must has unique
// value. In case if value is already exist in the skiplist node contains that
// value will be return along with false as last parameter
//
// If node has been inserted then function returns node itself with true
func (s *SkipList) Insert(node Node) (Node, bool) {
	inserted, ivalue, update := false, node.Value(), &linkedlist.State{}
	for !inserted {
		ipoint, t := s.Search(ivalue)
		if s.Compare(ivalue, t) == 0 {
			return t, false
		}

		// Insert onto data layer
		_, inserted = linkedlist.WeakInsert(ipoint, t, update, node)
	}

	// build index
	down, height := (*IndexNode)(nil), s.RandomLevel()
	for level := 1; level <= height && !linkedlist.LoadState(node).IsRemoved(); level++ {
		// Create "level" index layer node
		idx := &IndexNode{
			state: &linkedlist.State{},
			Root:  node,
			Down:  down,
		}

		inserted, update := false, &linkedlist.State{}
		for !inserted {
			// Find insertion point & insert
			ipoint, t := s.SearchIndex(ivalue, level)
			_, inserted = linkedlist.WeakInsert(ipoint, t, update, idx)
		}

		// next layer pls
		down = idx
	}

	return node, true
}

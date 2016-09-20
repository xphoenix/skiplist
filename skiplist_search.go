package skiplist

import "github.com/xphoenix/linkedlist"

// Search is looking for the two consequent nodes, so that first
// if strictly less then given key and second is greater or equals to the key.
//
// Note that function will always found a such pair because s.Head and s.Tail nodes
// are smaller/greater then any user supplied key by definition. Because of that
// returned nodes should be carefully checked by calling code
func (s *SkipList) Search(key interface{}) (Node, Node) {
	i1, _ := s.SearchIndex(key, 1)
	n1, n2 := i1.Root, linkedlist.LoadState(i1.Root).Next.(Node)

	for s.Compare(key, n2) > 0 {
		n1, n2 = n2, linkedlist.Next(n2).(Node)
	}

	return n1, n2
}

// SearchIndex is looking for the two consequent index nodes, so that first
// if strictly less then given key and second is greater or equals to the key
//
// Index node value is the value of the data node connected to the tower and all
// key comparisions gets done by the user supplied comparator defined in the
// skiplist
//
// tillLevel parameter tells what is the lowest index level should be considered
// by the algorithm
func (s *SkipList) SearchIndex(key interface{}, tillLevel int) (*IndexNode, *IndexNode) {
	level, cur, next := s.Height-1, s.IndexHead, linkedlist.Next(s.IndexHead).(*IndexNode)
	for cur.Down != nil && level > tillLevel {
		// fmt.Printf("[%d] %p -> %p\n", level, cur, next)

		// Note that Next above has alredy took care of the case if level node has been
		// concurrently deleted, however we must take care of the case if data node
		// for next tower has been removed.
		if linkedlist.LoadState(next.Root).IsRemoved() {
			// fmt.Printf("[%d] Removed\n", level)
			// If tower is superflovous - delete level node and restart Search
			// from the last valid node on the same level
			p, _, _ := linkedlist.WeakDelete(cur, next, &linkedlist.State{})
			cur = p.(*IndexNode)
		} else if s.Compare(key, next) <= 0 {
			// fmt.Printf("[%d] Down\n", level)
			// Can't move further on the current level as next node has too big
			// value, try lower level with smaller precission
			level, cur = level-1, cur.Down
		} else {
			// fmt.Printf("[%d] Next\n", level)
			// Next value on the same level is still too small, so lets
			// move further
			cur = next
		}

		// Read next node, that is the point determinates if search must be stop.
		next = linkedlist.Next(cur).(*IndexNode)
	}

	// Algorithm above found {cur, next} on the desired level, however
	// it might be still necessary to continue search on that last level
	for s.Compare(key, next) > 0 {
		// fmt.Printf("[%d] %p -> %p\n", level, cur, next)
		if linkedlist.LoadState(next.Root).IsRemoved() {
			// fmt.Printf("[%d] Removed\n", level)
			// If tower is superflovous - delete level node and restart Search
			// from the last valid node on the same level
			p, _, _ := linkedlist.WeakDelete(cur, next, &linkedlist.State{})
			cur, next = p.(*IndexNode), linkedlist.Next(p).(*IndexNode)
		} else {
			// fmt.Printf("[%d] Next\n", level)
			cur, next = next, linkedlist.Next(next).(*IndexNode)
		}
	}

	// done
	return cur, next
}

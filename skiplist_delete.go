package skiplist

import "github.com/xphoenix/linkedlist"

// Delete removed give value from the skiplist. If value found and was sucessfully
// removed then removed node returns otherwise nil
//
// Function returns two boolean flags. First one indicates that node has been found
// and removed, second flags is true only if node has been removed by the current
// thread.
//
// Latest flags needs if two concurrent go routines remove same node. In that case
// both will report successfull node deletition
func (s *SkipList) Delete(key interface{}) (Node, bool, bool) {
	update := &linkedlist.State{}
	for {
		prev, node := s.Search(key)
		if s.Compare(key, node) != 0 {
			return nil, false, false
		}

		_, removed, byme := linkedlist.WeakDelete(prev, node, update)
		if removed {
			// SearchIndex will scan each level, so that current superflovous tower
			// will be completely removed
			s.SearchIndex(key, 1)
			return node, removed, byme
		}
	}
}

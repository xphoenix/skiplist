package test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	. "github.com/xphoenix/skiplist"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Random tests
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type report struct {
	added      []int
	removed    []int
	duplicates int
}

func (r *report) merge(o *report) {
	r.added = append(r.added, o.added...)
	r.removed = append(r.removed, o.removed...)
	r.duplicates += o.duplicates
}

func (r *report) String() string {
	return fmt.Sprintf("added=%d, removed=%d, duplicates=%d", len(r.added), len(r.removed), r.duplicates)
}

func GenerateRangeChanges(sk *SkipList, base, size int, out chan<- *report) {
	info := &report{
		added:      make([]int, base),
		removed:    make([]int, base),
		duplicates: 0,
	}

	for i := 0; i < size; i++ {
		value := base - rand.Intn(base<<1)
		if value >= 0 {
			update := NewIntNode(value)
			n, inserted := sk.Insert(update)
			if inserted {
				info.added = append(info.added, value)
			} else if n != nil {
				info.duplicates++
			}
		} else {
			_, _, suc := sk.Delete(value)
			if suc {
				info.removed = append(info.removed, value)
			}
		}
	}
	out <- info
}

func TestRandomOperations(t *testing.T) {
	// That is quite time consuming test, so skip it in the short mode
	if testing.Short() {
		t.Skip("Skip test in short mode")
		return
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Setup
	// assert := assert.New(t)
	sk := New(IntNodeComparator, 8)

	// settings
	concurrency, base, out := 8, 100, make(chan *report)
	for i := 0; i < concurrency; i++ {
		go GenerateRangeChanges(sk, base, 1000*base, out)
	}

	// Wait for all goroutines done inserting data
	status, dc := &report{
		added:      make([]int, base),
		removed:    make([]int, base),
		duplicates: 0,
	}, 0
	for r := range out {
		status.merge(r)

		dc++
		if dc == concurrency {
			close(out)
		}
	}
	dump("TestRandomOperations.dot", sk)

	// Prepare data for validation
	// sort.Ints(status.added)
	// sort.Ints(status.removed)
	// column := make([]*IndexNode, sk.Height-1)
	// for p, l := sk.IndexHead, sk.Height-2; l >= 0; l-- {
	// 	column[l] = p
	// 	p = p.Down
	// }
	//
	// // Validate skiplist structure
	// dpointer := sk.DataHead
	// for column[0] != nil {
	// 	// Lower part of index might belong to cur dpointer column
	// 	ppoint := 0
	// 	for ; ppoint < sk.Height-1 && column[ppoint].Root == dpointer; ppoint++ {
	// 		column[ppoint] = linkedlist.LoadState(column[ppoint]).Next.(*IndexNode)
	// 	}
	//
	// 	// Tower could be lower then whole index but then rest of nodes belog to NEXT data node
	// 	for ; ppoint < sk.Height-1; ppoint++ {
	// 		assert.True(column[ppoint].Value().(int) > dpointer.Value().(int), "Index is straightforward")
	// 	}
	//
	// 	// TODO: check if value should be present (added, removed) arrays
	//
	// 	// scroll value
	// 	dpointer = linkedlist.LoadState(dpointer).Next.(Node)
	// }
}

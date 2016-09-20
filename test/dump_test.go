package test

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/xphoenix/linkedlist"
	"github.com/xphoenix/skiplist"
)

var graph bool

func init() {
	flag.BoolVar(&graph, "graph", false, "generate graphviz files for test lists")
	flag.Parse()
}

func dump(name string, sk *skiplist.SkipList) {
	if !graph {
		return
	}
	DumpSkiplist(name, sk)
}

// DumpSkiplist generate graphviz file with the skiplist content. It could be
// very useful for debugging
func DumpSkiplist(filename string, s *skiplist.SkipList) {
	// open file
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to create file: %s", err))
	}
	defer file.Close()

	// create writer
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Prepare column
	column := make([]linkedlist.Node, s.Height, s.Height)
	column[s.Height-1] = s.IndexHead
	for i := s.Height - 1; i > 0; i-- {
		column[i-1] = column[i].(*skiplist.IndexNode).Down
	}
	column[0] = s.DataHead

	// Start structure dump
	writer.WriteString("strict digraph skiplist {\n")
	writer.WriteString("\trank=same;\n")
	writer.WriteString("\trankdir=LR;\n\n")
	writer.WriteString("\tsplines=ortho;\n\n")

	// Dump nodes in columns
	for x := 0; column[0] != nil; x++ {
		writer.WriteString(fmt.Sprintf("\tsubgraph \"column#%d\"\n\t{\n", x))
		writer.WriteString(fmt.Sprintf("\t\tlabel=\"column %d\";\n", x))

		// Index layers
		var localHeight int
		for localHeight = 1; localHeight < s.Height && column[localHeight] != nil; localHeight++ {
			idx := column[localHeight].(*skiplist.IndexNode)
			if idx.Root != column[0] {
				break
			} else if idx.Down != nil {
				dumpEdge(writer, 50, column[localHeight], idx.Down)
			} else if idx.Root != nil {
				dumpEdge(writer, 0, column[localHeight], idx.Root)
			}
			dumpNode(writer, localHeight, idx)
		}
		// Data layer
		dumpNode(writer, 0, column[0])
		writer.WriteString("\t}\n\n")

		// Scroll columns
		for i := 0; i < localHeight; i++ {
			next := linkedlist.LoadState(column[i]).Next
			dumpEdge(writer, 100, column[i], next)
			column[i] = next
		}
	}

	writer.WriteString("}")
}

func dumpNode(out *bufio.Writer, level int, n linkedlist.Node) {
	if n == nil {
		return
	}

	id, s := nodeID(n), linkedlist.LoadState(n)

	color, style := "black", "solid"
	if s.IsFreezed() {
		color, style = "blue", "dashed"
	} else if s.IsRemoved() {
		color, style = "red", "dashed"
	}
	out.WriteString(fmt.Sprintf("\t\t%s [shape=record, style=%s, color=%s, label=\"{%d|%s}|{%s}\"];\n", id, style, color, level, id, n))
}

func dumpEdge(out *bufio.Writer, w int, src, dst linkedlist.Node) {
	if src == nil || dst == nil {
		return
	}

	l, r := nodeID(src), nodeID(dst)
	out.WriteString(fmt.Sprintf("\t%s -> %s [weight=%d];\n", l, r, w))
}

func nodeID(n linkedlist.Node) string {
	t, ok := n.(*skiplist.IndexNode)
	if ok {
		return fmt.Sprintf("node_%p", t)
	}

	t2, ok := n.(*skiplist.StabNode)
	if ok {
		return fmt.Sprintf("node_%p", t2)
	}

	t3, ok := n.(*IntNode)
	if ok {
		return fmt.Sprintf("node_%p", t3)
	}

	panic("Unknown node type")
}

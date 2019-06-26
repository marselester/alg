// Package digraph (directed graph) provides data structures to manipulate a set of vertices and
// a collection of directed edges. Each directed edge connects an ordered pair of vertices.
package digraph

import (
	"fmt"
	"strings"
)

// AdjacencyList is a digraph representation that maintains a vertex-indexed array of linked lists
// of the vertices adjacent to each vertex (each edge occurs just once, unlike undirected graph).
type AdjacencyList struct {
	a []*node
}
type node struct {
	// v is a vertex adjacent to the array-indexed vertex.
	v    int
	next *node
}

// NewAdjacencyList creates a graph represented as an array of adjacency lists.
func NewAdjacencyList(vertices int) *AdjacencyList {
	return &AdjacencyList{
		a: make([]*node, vertices),
	}
}

// Add adds a new edge connecting v and w vertices.
func (g *AdjacencyList) Add(v, w int) {
	g.a[v] = &node{v: w, next: g.a[v]}
}

// String returns a string representation of the graph's adjacency lists.
func (g *AdjacencyList) String() string {
	var b strings.Builder
	for i := 0; i < len(g.a); i++ {
		fmt.Fprintf(&b, "%d:", i)
		for n := g.a[i]; n != nil; n = n.next {
			fmt.Fprintf(&b, " %d", n.v)
		}
		fmt.Fprint(&b, "\n")
	}
	return b.String()
}

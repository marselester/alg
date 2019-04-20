package graph

import (
	"fmt"
	"strings"
)

/*
AdjacencyList maintains a vertex-indexed array of lists of the vertices adjacent to each vertex.
Every edge appears twice: if an edge connects v and w, then w appears in v's list
and v appears in w's list.
Space usage is proportional to number of vertices + edges.
A new edge is added in constant time and iteration through adjacent vertices is constant time
per adjacent vertex.
*/
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
	g.a[w] = &node{v: v, next: g.a[w]}
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

package spt

import "fmt"

// Edge represents digraph's edge.
type Edge struct {
	// V is "from" vertex.
	V int
	// W is "to" vertex.
	W      int
	Weight float64
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d->%d %.2f", e.V, e.W, e.Weight)
}

// AdjacencyList maintains a vertex-indexed array of adjacency lists of edges.
type AdjacencyList struct {
	a         []*node
	edgeCount int
}
type node struct {
	// edge is a vertex adjacent to the array-indexed vertex.
	edge *Edge
	next *node
}

// NewAdjacencyList creates a digraph represented as an array of adjacency lists.
func NewAdjacencyList(vertices int) *AdjacencyList {
	return &AdjacencyList{
		a: make([]*node, vertices),
	}
}

// Add adds a new edge connecting v and w vertices.
func (g *AdjacencyList) Add(e *Edge) {
	g.a[e.V] = &node{edge: e, next: g.a[e.V]}
	g.edgeCount++
}

// Edges returns all edges in this graph.
func (g *AdjacencyList) Edges() []*Edge {
	var edges []*Edge
	for v := 0; v < g.VertexCount(); v++ {
		for _, e := range g.Adjacent(v) {
			edges = append(edges, e)
		}
	}
	return edges
}

// Adjacent returns edges leaving vertex v.
func (g *AdjacencyList) Adjacent(v int) []*Edge {
	if v < 0 || v >= g.VertexCount() {
		return nil
	}

	var edges []*Edge
	for n := g.a[v]; n != nil; n = n.next {
		edges = append(edges, n.edge)
	}
	return edges
}

// VertexCount returns number of vertices in the digraph.
func (g *AdjacencyList) VertexCount() int {
	return len(g.a)
}

// EdgeCount returns number of edges in the digraph.
func (g *AdjacencyList) EdgeCount() int {
	return g.edgeCount
}

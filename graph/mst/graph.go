package mst

import "fmt"

// Edge represents graph's edge.
type Edge struct {
	V      int
	W      int
	Weight float32
}

// Either returns either of this edge's vertices.
// This is useful when neither vertex is known, a client can use the idiomatic code
// v, w = e.Either(), e.Other(v) to access an edge's two vertices.
func (e *Edge) Either() int {
	return e.V
}

// Other helps to fund the other vertex when it knows v.
// It returns -1 if v doesn't belong to the edge.
func (e *Edge) Other(v int) int {
	switch v {
	case e.V:
		return e.W
	case e.W:
		return e.V
	default:
		return -1
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.5f", e.V, e.W, e.Weight)
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

// NewAdjacencyList creates a graph represented as an array of adjacency lists.
func NewAdjacencyList(vertices int) *AdjacencyList {
	return &AdjacencyList{
		a: make([]*node, vertices),
	}
}

// Add adds a new edge connecting v and w vertices.
func (g *AdjacencyList) Add(e *Edge) {
	v := e.Either()
	w := e.Other(v)
	g.a[v] = &node{edge: e, next: g.a[v]}
	g.a[w] = &node{edge: e, next: g.a[w]}
	g.edgeCount++
}

// Edges returns all edges in this graph.
func (g *AdjacencyList) Edges() []*Edge {
	var edges []*Edge
	for v := 0; v < g.VertexCount(); v++ {
		for _, e := range g.Adjacent(v) {
			if e.Other(v) > v {
				edges = append(edges, e)
			}
		}
	}
	return edges
}

// Adjacent returns edges incident to vertex v.
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

// VertexCount returns number of vertices in the graph.
func (g *AdjacencyList) VertexCount() int {
	return len(g.a)
}

// EdgeCount returns number of edges in the graph.
func (g *AdjacencyList) EdgeCount() int {
	return g.edgeCount
}

package graph

import (
	"bufio"
	"fmt"
	"io"
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

// Adjacent returns vertices adjacent to the vertex v.
func (g *AdjacencyList) Adjacent(v int) []int {
	var edges []int
	if v < 0 || v >= len(g.a) {
		return edges
	}
	for n := g.a[v]; n != nil; n = n.next {
		edges = append(edges, n.v)
	}
	return edges
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

// SymbolGraph is a graph whose vertices are strings (city, movie), not integer indices.
type SymbolGraph struct {
	// names maps a vertex name (city, movie) to its index in the graph.
	names map[string]int
	// keys is an inverted index of vertex names.
	keys []string
	g    *AdjacencyList
}

// NewSymbolGraph constructs a symbol graph. It uses two passes through the data to build
// underlying indices. Note, one pass would be sufficient if a graph used a hash table
// implementation (extra log V factor, V is number of vertices) instead of a linked list.
func NewSymbolGraph(in io.ReadSeeker, sep string) (*SymbolGraph, error) {
	sg := SymbolGraph{
		names: make(map[string]int),
	}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		// Edges connect the first vertex name to other vertices, e.g., movie and performers.
		edges := strings.Split(scanner.Text(), sep)
		for _, name := range edges {
			if _, ok := sg.names[name]; !ok {
				sg.names[name] = len(sg.names)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	sg.keys = make([]string, len(sg.names))
	for name, index := range sg.names {
		sg.keys[index] = name
	}

	sg.g = NewAdjacencyList(len(sg.names))
	if _, err := in.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	scanner = bufio.NewScanner(in)
	for scanner.Scan() {
		edges := strings.Split(scanner.Text(), sep)
		src, _ := sg.Index(edges[0])
		for i := 1; i < len(edges); i++ {
			dst, _ := sg.Index(edges[i])
			sg.g.Add(src, dst)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &sg, nil
}

// Adjacent returns vertices adjacent to the vertex v.
func (sg *SymbolGraph) Adjacent(v int) []int {
	return sg.g.Adjacent(v)
}

// Index returns a vertex index associated with a key (vertex name).
func (sg *SymbolGraph) Index(key string) (int, bool) {
	v, ok := sg.names[key]
	return v, ok
}

// Name returns a vertex name associated with index v.
// Empty string indicates that vertex doesn't exist.
func (sg *SymbolGraph) Name(v int) string {
	if v < 0 || v >= len(sg.keys) {
		return ""
	}
	return sg.keys[v]
}

// Graph returns an underlying graph.
func (sg *SymbolGraph) Graph() *AdjacencyList {
	return sg.g
}

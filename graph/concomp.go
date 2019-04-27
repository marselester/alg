package graph

/*
NewConnectedComponent finds the connected components of a graph.
It uses marked array to find a vertex to serve as the starting point for a depth-first search
in each component. Firstly it marks all vertices connected to 0. Then in a loop it looks for
an unmarked vertex and marks all vertices connected to that vertex.

DFS uses preprocessing time and space proportional to E+V (number of edges plus vertices) to support
constant-time connectivity queries in a graph. In theory DFS is faster than union-find algorithm,
because it provides a constant-time guarantee. In practise, the difference is negligible,
and union-find is faster because it doesn't have to build a full representation of the graph.
*/
func NewConnectedComponent(g *AdjacencyList) *ConnectedComponent {
	cc := ConnectedComponent{
		g:      g,
		marked: make([]bool, len(g.a)),
		sites:  make([]int, len(g.a)),
	}
	// It finds an unmarked vertex and calls the recursive search() to mark and identify
	// all vertices connected to it, continuing until all vertices have been marked and identified.
	for v := 0; v < len(cc.sites); v++ {
		if !cc.marked[v] {
			cc.search(v)
			cc.count++
		}
	}
	return &cc
}

// ConnectedComponent uses depth-first search to find connected components in a graph.
type ConnectedComponent struct {
	g *AdjacencyList
	// marked is an array of visited vertices.
	marked []bool
	// sites is an array of sites (e.g., a computer in a network).
	// An index represents a site number (vertex), value corresponds its component ID.
	// In other words, where a site belongs to.
	sites []int
	// count is a number of components in the network.
	count int
}

// search visits all vertices connected to the source vertex. To visit a vertex
// mark it as visited and visit (recursively) all the vertices that are adjacent to it and
// that have not been marked yet.
func (cc *ConnectedComponent) search(source int) {
	cc.marked[source] = true
	// Vertices connected to the source belong to the same component.
	cc.sites[source] = cc.count
	for n := cc.g.a[source]; n != nil; n = n.next {
		if !cc.marked[n.v] {
			cc.search(n.v)
		}
	}
}

// IsConnected tells whether v and w are in the same component
// by checking if site identifiers are equal.
func (cc *ConnectedComponent) IsConnected(v, w int) bool {
	return cc.sites[v] == cc.sites[w]
}

// Count returns number of components.
func (cc *ConnectedComponent) Count() int {
	return cc.count
}

// ID returns component identifier (between 0 and count-1) for vertex v.
func (cc *ConnectedComponent) ID(v int) int {
	return cc.sites[v]
}

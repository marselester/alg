package digraph

/*
NewStrongConnectedComponent finds strong connected components of a digraph using Kosaraju-Sharir algorithm.

Strong connectivity helps to understand the structure of a digraph,
highlighting interrelated sets of vertices (strong components).
For example, you can model a food chain (predator-prey) to study ecological system.
Another example is web content where a page represents a vertex, and edge represents a hyperlink.

Two vertices v and w are strongly connected if there is a directed path from v to w,
and a directed path from w to v (if there exists a general directed cycle that contains them both).
A digraph is strongly connected if all its vertices are strongly connected to one another.
*/
func NewStrongConnectedComponent(g *AdjacencyList) *StrongConnectedComponent {
	scc := StrongConnectedComponent{
		g:      g,
		marked: make([]bool, g.VertexCount()),
		sites:  make([]int, g.VertexCount()),
	}

	for _, v := range reversePostorder(Reverse(g)) {
		if !scc.marked[v] {
			scc.dfs(v)
			scc.count++
		}
	}
	return &scc
}

/*
StrongConnectedComponent helps to check if two given vertices are strongly connected and
how many strong components the digraph has.

It uses depth-first search to find connected components in a digraph as follows:

	Compute reverse topological order: the reverse postorder of the reversed digraph.
	Run DFS and consider the unmarked vertices in the order computed on the previous step.
	All vertices visited on a call to the recursive dfs() from NewConnectedComponent are strong components.

The Kosaraju-Sharir algorithm uses preprocessing time and space proportional to E+V to support constant-time
strong connectivity queries in a digraph. It computes the reverse of the digraph and does two depth-first searches.
Each step takes time proportional to E+V. The reverse copy of the digraph uses space proportional to E+V.
*/
type StrongConnectedComponent struct {
	g *AdjacencyList
	// marked is an array of visited vertices.
	marked []bool
	// sites is an array of sites (e.g., a computer in a network).
	// An index represents a site number (vertex), value corresponds its component ID.
	// In other words, where a site belongs to.
	sites []int
	// count is a number of strong components in the network.
	count int
}

func (scc *StrongConnectedComponent) dfs(source int) {
	scc.marked[source] = true
	scc.sites[source] = scc.count
	for n := scc.g.a[source]; n != nil; n = n.next {
		if !scc.marked[n.v] {
			scc.dfs(n.v)
		}
	}
}

// IsConnected tells whether v and w are in the same component
// by checking if site identifiers are equal.
func (scc *StrongConnectedComponent) IsConnected(v, w int) bool {
	return scc.sites[v] == scc.sites[w]
}

// Count returns number of components.
func (scc *StrongConnectedComponent) Count() int {
	return scc.count
}

// ID returns component identifier (between 0 and count-1) for vertex v.
func (scc *StrongConnectedComponent) ID(v int) int {
	return scc.sites[v]
}

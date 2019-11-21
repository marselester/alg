package digraph

// TransitiveClosure solves all-pairs reachability problem (is there a path from vertex v to w?)
// for small or dense digraphs.
// A general solution that achieves constant-time queries with substantially less than
// quadratic space is an unsolved research problem (no practical solution for giant digraph such as the web graph).
//
// The transitive closure of a digraph G is another digraph with the same set of vertices,
// but with an edge from v to w if and only if w is reachable from v in G.
// Since transitive closures are typically dense, they are normally represented with a matrix of bool values,
// where the entry in row v and column w is true if w is reachable from v.
type TransitiveClosure struct {
	// matrix is a transitive closure matrix where v-th row represents
	// results of depth-first search for vertex v (vertices that are reachable from v).
	matrix [][]int
}

// NewTransitiveClosure constructs a transitive closure to check vertex reachability.
// It uses space proportional to V^2 (V is a number of vertices) and time proportional
// to V * (E + V) where E is a number of edges.
func NewTransitiveClosure(g *AdjacencyList) *TransitiveClosure {
	tc := TransitiveClosure{
		matrix: make([][]int, g.VertexCount()),
	}
	for v := 0; v < g.VertexCount(); v++ {
		tc.matrix[v] = DepthFirstSearch(g, v)
	}
	return &tc
}

// Reachable returns true if vertex w is reachable from v.
func (tc *TransitiveClosure) Reachable(v, w int) bool {
	// This loop is redundant in case marked array is accessible from DepthFirstSearch.
	// For example, code would be: return tc.matrix[v].marked(w).
	for i := range tc.matrix[v] {
		if tc.matrix[v][i] == w {
			return true
		}
	}
	return false
}

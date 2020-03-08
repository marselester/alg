package spt

import (
	"math"

	"github.com/marselester/alg/digraph/weighted"
)

// NewAcyclic finds shortest path in edge-weighted DAG:
// init distTo[source] to zero and all other distTo[] entries to positive infinity;
// then relax the vertices, one by one, taking the vertices in topological order.
// Algorithm solves the single-source problem in linear time (E + V), handles negative edge weights.
func NewAcyclic(g *weighted.AdjacencyList, source int) *Acyclic {
	a := Acyclic{
		g:      g,
		edgeTo: make([]*weighted.Edge, g.VertexCount()),
		distTo: make([]float64, g.VertexCount()),
	}

	// Start with distTo[source] = 0 and all other entries equal to positive infinity.
	for i := 0; i < len(a.distTo); i++ {
		a.distTo[i] = math.Inf(+1)
	}
	a.distTo[source] = 0

	for _, v := range weighted.TopologicalSort(g) {
		a.relax(v)
	}

	return &a
}

// Acyclic uses topological sort for acyclic edge-weighted digraphs to compute shortest paths.
type Acyclic struct {
	g *weighted.AdjacencyList
	// edgeTo is a parent-edge representation as vertex-indexed array
	// where edgeTo[v] is the edge that connects v to its parent in the tree
	// (the last edge on a shortest path from source to v).
	// It helps to find edges on the shortest-paths tree.
	// By convention, edgeTo[source] is nil.
	edgeTo []*weighted.Edge
	// distTo is vertex-indexed array such that distTo[v] is the length of the shortest known path from source to v.
	// It helps to find distance to the source.
	// By convention, distTo[source] is 0.
	distTo []float64
}

// DistTo returns distance from source vertex to v, or infinity if no path exists.
func (a *Acyclic) DistTo(v int) float64 {
	return a.distTo[v]
}

// HasPathTo returns true if there is a path from source vertex to v
// (whether v is reachable from source vertex) by checking whether distance to v is finite.
func (a *Acyclic) HasPathTo(v int) bool {
	return !math.IsInf(a.distTo[v], +1)
}

// PathTo returns a path from source vertex to v or nil if it doesn't exist.
func (a *Acyclic) PathTo(v int) []*weighted.Edge {
	if !a.HasPathTo(v) {
		return nil
	}
	return weighted.PathTo(v, a.edgeTo)
}

// relax relaxes all the edges leaving vertex v.
func (a *Acyclic) relax(v int) {
	for _, e := range a.g.Adjacent(v) {
		relax(e, a.distTo, a.edgeTo)
	}
}

// relax performs edge-relaxation operation (relaxing the tension on the rubber band along a shorter path)
// and reports if the edge was eligible.
//
// To relax an edge v->w means to test whether the best known way from source to w is to go from source to v,
// then take the edge from v to w.
// The best known distance to w through v is the sum of distTo[v] and edge weight:
//
//     if the sum is smaller than distTo[w], then edge v->w is eligible
//     otherwise v->w is ineligible (ignored)
//
// For example, length of the shortest path from source to v is distTo[e.V] = 3.1,
// length of the shortest path from source to w is distTo[e.W] = 3.3.
// If e.Weight is 1.3, then edge e is ignored, because 3.3 < 3.1 + 1.3.
// If weight is 0.1, then edge e leads to a shorter path to w, because 3.3 > 3.1 + 0.1.
func relax(e *weighted.Edge, distTo []float64, edgeTo []*weighted.Edge) bool {
	distance := distTo[e.V] + e.Weight
	isEligible := distTo[e.W] > distance
	if isEligible {
		distTo[e.W] = distance
		edgeTo[e.W] = e
	}
	return isEligible
}

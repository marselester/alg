package lpt

import (
	"math"

	"github.com/marselester/alg/digraph/weighted"
)

// NewAcyclic finds longest path in edge-weighted DAG:
// init distTo[source] to zero and all other distTo[] entries to negative infinity;
// then relax the vertices, one by one, taking the vertices in topological order.
// Algorithm solves the single-source problem in linear time (E + V), handles negative edge weights.
func NewAcyclic(g *weighted.AdjacencyList, source int) *Acyclic {
	a := Acyclic{
		g:      g,
		edgeTo: make([]*weighted.Edge, g.VertexCount()),
		distTo: make([]float64, g.VertexCount()),
	}

	// Start with distTo[source] = 0 and all other entries equal to negative infinity.
	for i := 0; i < len(a.distTo); i++ {
		a.distTo[i] = math.Inf(-1)
	}
	a.distTo[source] = 0

	for _, v := range weighted.TopologicalSort(g) {
		a.relax(v)
	}

	return &a
}

// Acyclic uses topological sort for acyclic edge-weighted digraphs to compute longest paths.
type Acyclic struct {
	g      *weighted.AdjacencyList
	edgeTo []*weighted.Edge
	distTo []float64
}

// relax relaxes all the edges leaving vertex v.
func (a *Acyclic) relax(v int) {
	for _, e := range a.g.Adjacent(v) {
		relax(e, a.distTo, a.edgeTo)
	}
}

// relax performs edge-relaxation operation just like in spt package, but with different sense of inequality.
func relax(e *weighted.Edge, distTo []float64, edgeTo []*weighted.Edge) bool {
	distance := distTo[e.V] + e.Weight
	isEligible := distTo[e.W] < distance
	if isEligible {
		distTo[e.W] = distance
		edgeTo[e.W] = e
	}
	return isEligible
}

// DistTo returns distance from source vertex to v, or infinity if no path exists.
func (a *Acyclic) DistTo(v int) float64 {
	return a.distTo[v]
}

// HasPathTo returns true if there is a path from source vertex to v
// (whether v is reachable from source vertex) by checking whether distance to v is finite.
func (a *Acyclic) HasPathTo(v int) bool {
	return !math.IsInf(a.distTo[v], -1)
}

// PathTo returns a path from source vertex to v or nil if it doesn't exist.
func (a *Acyclic) PathTo(v int) []*weighted.Edge {
	if !a.HasPathTo(v) {
		return nil
	}
	return weighted.PathTo(v, a.edgeTo)
}

package spt

import (
	"math"
)

// NewAcyclic finds shortest path in edge-weigthed DAG:
// init distTo[source] to zero and all other distTo[] entries to positive infinity;
// then relax the vertices, one by one, taking the vertices in topological order.
// Algorithm solves the single-source problem in linear time (E + V), handles negative edge weights.
func NewAcyclic(g *AdjacencyList, source int) *Acyclic {
	a := Acyclic{
		g:      g,
		edgeTo: make([]*Edge, g.VertexCount()),
		distTo: make([]float64, g.VertexCount()),
	}

	// Start with distTo[source] = 0 and all other entries equal to positive infinity.
	for i := 0; i < len(a.distTo); i++ {
		a.distTo[i] = math.Inf(+1)
	}
	a.distTo[source] = 0

	for _, v := range topologicalSort(g) {
		a.relax(v)
	}

	return &a
}

// Acyclic uses topological sort for acyclic edge-weighted digraphs to compute shortest paths.
type Acyclic struct{
	g *AdjacencyList
	// edgeTo is a parent-edge representation as vertex-indexed array
	// where edgeTo[v] is the edge that connects v to its parent in the tree
	// (the last edge on a shortest path from source to v).
	// It helps to find edges on the shortest-paths tree.
	// By convention, edgeTo[source] is nil.
	edgeTo []*Edge
	// distTo is vertex-indexed array such that distTo[v] is the length of the shortest known path from source to v.
	// It helps to find distance to the source.
	// By convention, distTo[source] is 0.
	distTo []float64
}

// relax relaxes all the edges leaving vertex v.
func (a *Acyclic) relax(v int) {
	for _, e := range a.g.Adjacent(v) {
		relax(e, a.distTo, a.edgeTo)
	}
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
func (a *Acyclic) PathTo(v int) []*Edge {
	if !a.HasPathTo(v) {
		return nil
	}
	return pathTo(v, a.edgeTo)
}

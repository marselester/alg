package spt

import (
	"math"

	"github.com/marselester/alg/sort/pqueue"
)

// NewDijkstra finds shortest paths using Dijkstra's algorithm:
// init distTo[source] to zero and all other distTo[] entries to positive infinity;
// then relax and add to the shortest-path tree (SPT) a non-tree vertex with the lowest distTo[] value,
// continuing until all vertices are on the tree or no non-tree vertex has a finite distTo[] value.
//
// Dijkstra's algorithm uses extra space proportional to V and time proportional to E * log V (in worst case).
// Another way to think about Dijkstra's algorithm is to compare it to eager version of Prim's algorithm.
func NewDijkstra(g *AdjacencyList, source int) *Dijkstra {
	d := Dijkstra{
		g:      g,
		edgeTo: make([]*Edge, g.VertexCount()),
		distTo: make([]float64, g.VertexCount()),
		pq:     pqueue.NewIndexMinHeap(g.VertexCount()),
	}

	// Start with distTo[source] = 0 and all other entries equal to positive infinity.
	for i := 1; i < len(d.distTo); i++ {
		d.distTo[i] = math.Inf(+1)
	}

	d.pq.Insert(source, 0)
	for d.pq.Size() != 0 {
		v, _ := d.pq.Min()
		d.relax(v)
	}

	return &d
}

// Dijkstra is an implementation of Dijkstra's algorithm which solves the single-source shortest-paths problem
// in edge-weighted digraphs with nonnegative weights.
type Dijkstra struct {
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
	// pq keeps track of vertices that are candidates for being the next to be relaxed.
	pq *pqueue.IndexMinHeap
}

// relax relaxes all the edges leaving vertex v.
// Vertex e.W is added to PQ if it's not there yet.
// Otherwise its priority is lowered.
func (d *Dijkstra) relax(v int) {
	for _, e := range d.g.Adjacent(v) {
		if relax(e, d.distTo, d.edgeTo) {
			if d.pq.Contains(e.W) {
				d.pq.Update(e.W, d.distTo[e.W])
			} else {
				d.pq.Insert(e.W, d.distTo[e.W])
			}
		}
	}
}

// DistTo returns distance from source vertex to v, or infinity if no path exists.
func (d *Dijkstra) DistTo(v int) float64 {
	return d.distTo[v]
}

// HasPathTo returns true if there is a path from source vertex to v
// (whether v is reachable from source vertex) by checking whether distance to v is finite.
func (d *Dijkstra) HasPathTo(v int) bool {
	return !math.IsInf(d.distTo[v], +1)
}

// PathTo returns a path from source vertex to v or nil if it doesn't exist.
func (d *Dijkstra) PathTo(v int) []*Edge {
	if !d.HasPathTo(v) {
		return nil
	}
	return pathTo(v, d.edgeTo)
}

// NewDijkstraAllPairs finds shortest paths for all vertex pairs
// using time proportional to E * V * log V and extra space proportional to V^2.
// It builds an array of Dijkstra objects, one for each vertex as the source.
// To find a shortest path, it uses the source to access the corresponding single-source
// shortest-paths object and then passes the target as an argument to the query.
func NewDijkstraAllPairs(g *AdjacencyList) *DijkstraAllPairs {
	all := DijkstraAllPairs{
		sources: make([]*Dijkstra, g.VertexCount()),
	}
	for v := 0; v < g.VertexCount(); v++ {
		all.sources[v] = NewDijkstra(g, v)
	}
	return &all
}

// DijkstraAllPairs solves all-pairs shortest paths problem.
type DijkstraAllPairs struct {
	sources []*Dijkstra
}

// PathTo returns a path from source to target vertex or nil if it doesn't exist.
func (all *DijkstraAllPairs) PathTo(source, target int) []*Edge {
	return all.sources[source].PathTo(target)
}

// DistTo returns distance from source to target vertex, or infinity if no path exists.
func (all *DijkstraAllPairs) DistTo(source, target int) float64 {
	return all.sources[source].DistTo(target)
}

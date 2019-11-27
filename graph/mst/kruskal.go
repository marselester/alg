package mst

import "github.com/marselester/alg/unionfind/wqunion"

// NewKruskal processes the edges in order of their weight (smallest to largest),
// taking for the MST each edge that doesn't form a cycle with edges previously added,
// stopping after adding V-1 edges.
// The edges form a forest of trees that evolves gradually into a single tree, the MST.
//
// Kruskal's algorithm uses space proportional to E and time proportional to E * log E in the worst case.
// The cost bound is conservative, since the algorithm terminates after finding V-1 MST edges.
// Kruskal's algorithm is generally slower than Prim's because it has to do a connected operation for each edge.
func NewKruskal(g *AdjacencyList) *Kruskal {
	k := Kruskal{
		g:  g,
		pq: NewMinHeap(g.EdgeCount()),
		uf: wqunion.New(g.VertexCount()),
	}

	for _, e := range g.Edges() {
		k.pq.Insert(e)
	}

	for k.pq.Size() != 0 && len(k.edges) < g.VertexCount()-1 {
		// Get lowest-weight edge from priority queue.
		e := k.pq.Min()
		v := e.Either()
		w := e.Other(v)

		// Skip if ineligible.
		if k.uf.IsConnected(v, w) {
			continue
		}
		// Merge components.
		k.uf.Connect(v, w)

		// Add edge to tree.
		k.edges = append(k.edges, e)
	}

	return &k
}

// Kruskal is an implementation of Kruskal's algorithm which uses priority queue
// to hold edges not yet examined, and a union-find data structure for identifying ineligible edges.
type Kruskal struct {
	g *AdjacencyList
	// pq is min priority queue to consider the edges in order by weight.
	pq *MinHeap
	// uf is a union-find to identify edges that cause cycles.
	uf *wqunion.Network
	// edges holds MST edges.
	edges []*Edge
}

// Edges returns all of the MST edges in increasing order of their weights.
func (k *Kruskal) Edges() []*Edge {
	return k.edges
}

// Weight calculates weight of MST.
// It requires iterating through the tree edges to add up the edge weights (lazy approach).
// In eager approach a running total would be kept.
func (k *Kruskal) Weight() float32 {
	var sum float32
	for i := range k.edges {
		sum += k.edges[i].Weight
	}
	return sum
}

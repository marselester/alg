package mst

// NewLazyPrim computes MST using Prim's algorithm:
// take an edge from the priority queue and (if it is eligible) add it to the tree,
// and also add to the tree the new vertex that it leads to,
// updating the set of crossing edges by calling visit with that vertex as argument.
//
// This implementation is a lazy approach where ineligible edges are left in the priority queue.
func NewLazyPrim(g *AdjacencyList) *LazyPrim {
	lp := LazyPrim{
		g:      g,
		marked: make([]bool, g.VertexCount()),
		pq:     NewMinHeap(g.EdgeCount()),
	}

	lp.visit(0)
	for lp.pq.Size() != 0 {
		// Get lowest-weight edge from priority queue.
		e := lp.pq.Min()
		v := e.Either()
		w := e.Other(v)

		// Skip if ineligible.
		if lp.marked[v] && lp.marked[w] {
			continue
		}

		// Add edge to tree.
		lp.edges = append(lp.edges, e)

		// Add vertex to tree (either v or w).
		if !lp.marked[v] {
			lp.visit(v)
		}
		if !lp.marked[w] {
			lp.visit(w)
		}
	}

	return &lp
}

// LazyPrim is a lazy version of Prim's algorithm to compute MST.
// It uses space proportional to E (number of edges) and time proportional to E * log E in the worst case.
// The bottleneck is the number of edge-weighted comparisons in the priority queue methods insert and delete min.
// In practice, the upper bound on the running time is a bit conservative because the number of edges
// on the priority queue is typically much less than E.
type LazyPrim struct {
	g *AdjacencyList
	// marked represents vertices on the tree. It is a vertex-indexed bool array
	// where marked[v] is true if v is on the tree.
	marked []bool
	// pq is min priority queue that compares crossing edges by weight to find the crossing edge of minimal weight.
	pq *MinHeap
	// edges holds MST edges.
	edges []*Edge
}

// visit puts a vertex on the tree, by marking it as visited and then
// putting all of its incident edges (eligible) onto the priority queue,
// thus ensuring the priority queue contains the crossing edges from tree vertices to non-tree vertices.
func (lp *LazyPrim) visit(v int) {
	lp.marked[v] = true
	for _, e := range lp.g.Adjacent(v) {
		if !lp.marked[e.Other(v)] {
			lp.pq.Insert(e)
		}
	}
}

// Edges returns all of the MST edges.
func (lp *LazyPrim) Edges() []*Edge {
	return lp.edges
}

// Weight calculates weight of MST.
// It requires iterating through the tree edges to add up the edge weights (lazy approach).
// In eager approach a running total would be kept.
func (lp *LazyPrim) Weight() float32 {
	var sum float32
	for i := range lp.edges {
		sum += lp.edges[i].Weight
	}
	return sum
}

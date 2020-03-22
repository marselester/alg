package spt

import (
	"math"

	"github.com/marselester/alg/digraph/weighted"
)

// NewBellmanFord finds shortest paths using Bellman-Ford algorithm.
func NewBellmanFord(g *weighted.AdjacencyList, source int) *BellmanFord {
	bf := BellmanFord{
		g:       g,
		edgeTo:  make([]*weighted.Edge, g.VertexCount()),
		distTo:  make([]float64, g.VertexCount()),
		onQueue: make([]bool, g.VertexCount()),
	}

	// Start with distTo[source] = 0 and all other entries equal to positive infinity.
	for i := 0; i < len(bf.distTo); i++ {
		bf.distTo[i] = math.Inf(+1)
	}
	bf.distTo[source] = 0

	// We start by putting the source onto the queue, then enter a loop where we take
	// a vertex off the queue and relax it.
	//
	// The digraph has a negative cycle reachable from the source if and only if the queue is non-empty
	// after the Vth pass through all the edges.
	bf.q.Enqueue(source)
	bf.onQueue[source] = true
	for bf.q.Size() > 0 && !bf.HasNegativeCycle() {
		v := bf.q.Dequeue()
		bf.onQueue[v] = false
		bf.relax(v)
	}

	return &bf
}

// BellmanFord is an implementation of queue-based Bellman-Ford algorithm
// which solves the single-source shortest-paths problem when negative cycles are not reachable.
// For example, in a job-scheduling-with-deadlines problem, negative cycles are rare due to
// logical real-world constraints.
// The algorithm takes time proportional to E * V and extra space proportional to V.
type BellmanFord struct {
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
	// onQueue is vertex-indexed array that tells whether vertex is on the queue, to avoid duplicates.
	onQueue []bool
	// q contanis vertices being relaxed.
	q queue
	// cost is a number of edge relaxations (number of passes).
	cost int
	// cycle is a first cycle (vertices from last known source vertex back to itself) found in the directed graph.
	cycle []int
}

// relax relaxes all the edges leaving vertex v.
// It puts vertices whose distTo value changes onto a FIFO queue (avoiding duplicates) and
// periodically checks for a negative cycle in edgeTo.
func (bf *BellmanFord) relax(v int) {
	for _, e := range bf.g.Adjacent(v) {
		if relax(e, bf.distTo, bf.edgeTo) {
			// Only one copy of each vertex appears on the queue.
			if bf.onQueue[e.W] {
				continue
			}
			bf.q.Enqueue(e.W)
			bf.onQueue[e.W] = true
		}

		// Ensure the algorithm terminates after V passes.
		bf.cost++
		if bf.cost%bf.g.VertexCount() == 0 {
			bf.findNegativeCycle()
		}
	}
}

// findNegativeCycle checks if edgeTo has a negative cycle to avoid an infinite loop in relax function.
func (bf *BellmanFord) findNegativeCycle() {
	g := weighted.NewAdjacencyList(bf.g.VertexCount())
	for v := 0; v < g.VertexCount(); v++ {
		if bf.edgeTo[v] != nil {
			g.Add(bf.edgeTo[v])
		}
	}
	bf.cycle = weighted.Cycle(g)
}

// HasNegativeCycle reports whether digraph has a negative cycle.
func (bf *BellmanFord) HasNegativeCycle() bool {
	return bf.cycle != nil
}

// NegativeCycle returns a negative cycle: directed cycle whose total weight
// (sum of the weights of its edges) is negative.
func (bf *BellmanFord) NegativeCycle() []int {
	return bf.cycle
}

// DistTo returns distance from source vertex to v, or infinity if no path exists.
func (bf *BellmanFord) DistTo(v int) float64 {
	return bf.distTo[v]
}

// HasPathTo returns true if there is a path from source vertex to v
// (whether v is reachable from source vertex) by checking whether distance to v is finite.
func (bf *BellmanFord) HasPathTo(v int) bool {
	return !math.IsInf(bf.distTo[v], +1)
}

// PathTo returns a path from source vertex to v or nil if it doesn't exist.
func (bf *BellmanFord) PathTo(v int) []*weighted.Edge {
	if !bf.HasPathTo(v) {
		return nil
	}
	return weighted.PathTo(v, bf.edgeTo)
}

// queue represents a queue of vertices.
type queue struct {
	vertices []int
}

// Enqueue adds a vertex to the end of the queue.
func (q *queue) Enqueue(v int) {
	q.vertices = append(q.vertices, v)
}

// Dequeue returns a vertex from the beginning of the queue.
func (q *queue) Dequeue() int {
	if len(q.vertices) == 0 {
		return 0
	}
	v := q.vertices[0]
	q.vertices = q.vertices[1:]
	return v
}

// Size returns the number of vertices in the queue.
func (q *queue) Size() int {
	return len(q.vertices)
}

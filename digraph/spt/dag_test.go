package spt

import (
	"fmt"
	"testing"

	"github.com/marselester/alg/digraph/weighted"
)

func ExampleNewAcyclic() {
	g := weighted.NewAdjacencyList(8)
	g.Add(&weighted.Edge{5, 4, 0.35})
	g.Add(&weighted.Edge{4, 7, 0.37})
	g.Add(&weighted.Edge{5, 7, 0.28})
	g.Add(&weighted.Edge{5, 1, 0.32})
	g.Add(&weighted.Edge{4, 0, 0.38})
	g.Add(&weighted.Edge{0, 2, 0.26})
	g.Add(&weighted.Edge{3, 7, 0.39})
	g.Add(&weighted.Edge{1, 3, 0.29})
	g.Add(&weighted.Edge{7, 2, 0.34})
	g.Add(&weighted.Edge{6, 2, 0.4})
	g.Add(&weighted.Edge{3, 6, 0.52})
	g.Add(&weighted.Edge{6, 0, 0.58})
	g.Add(&weighted.Edge{6, 4, 0.93})

	a := NewAcyclic(g, 5)
	for v := 0; v < g.VertexCount(); v++ {
		fmt.Printf("5 to %d: %v\n", v, a.PathTo(v))
	}
	// Output:
	// 5 to 0: [5->4 0.35 4->0 0.38]
	// 5 to 1: [5->1 0.32]
	// 5 to 2: [5->7 0.28 7->2 0.34]
	// 5 to 3: [5->1 0.32 1->3 0.29]
	// 5 to 4: [5->4 0.35]
	// 5 to 5: []
	// 5 to 6: [5->1 0.32 1->3 0.29 3->6 0.52]
	// 5 to 7: [5->7 0.28]
}

func TestRelax(t *testing.T) {
	e := &weighted.Edge{V: 0, W: 1}
	// Length of the shortest path from source to v is distTo[e.V] = 3.1,
	// length of the shortest path from source to w is distTo[e.W] = 3.3.
	distTo := make([]float64, 2, 2)
	distTo[e.V] = 3.1
	distTo[e.W] = 3.3

	// When weight is 0.1, edge e leads to a shorter path to w, because 3.3 > 3.1 + 0.1.
	e.Weight = 0.1
	edgeTo := make([]*weighted.Edge, 2, 2)
	if ok := relax(e, distTo, edgeTo); !ok {
		t.Errorf("relax() = false when dist to v + edge weight < dist to w")
	}
	if edgeTo[e.V] != nil || edgeTo[e.W] != e {
		t.Errorf("relax() edgeTo was not updated: %v", edgeTo)
	}

	// When e.Weight is 1.3, edge e is ignored, because 3.3 < 3.1 + 1.3.
	e.Weight = 1.3
	edgeTo = make([]*weighted.Edge, 2, 2)
	if ok := relax(e, distTo, edgeTo); ok {
		t.Errorf("relax() = true when dist to v + edge weight > dist to w")
	}
	if edgeTo[e.V] != nil || edgeTo[e.W] != nil {
		t.Errorf("relax() edgeTo was updated: %v", edgeTo)
	}
}

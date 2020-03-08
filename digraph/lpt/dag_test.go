package lpt

import (
	"fmt"

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
	// 5 to 0: [5->1 0.32 1->3 0.29 3->6 0.52 6->4 0.93 4->0 0.38]
	// 5 to 1: [5->1 0.32]
	// 5 to 2: [5->1 0.32 1->3 0.29 3->6 0.52 6->4 0.93 4->7 0.37 7->2 0.34]
	// 5 to 3: [5->1 0.32 1->3 0.29]
	// 5 to 4: [5->1 0.32 1->3 0.29 3->6 0.52 6->4 0.93]
	// 5 to 5: []
	// 5 to 6: [5->1 0.32 1->3 0.29 3->6 0.52]
	// 5 to 7: [5->1 0.32 1->3 0.29 3->6 0.52 6->4 0.93 4->7 0.37]
}

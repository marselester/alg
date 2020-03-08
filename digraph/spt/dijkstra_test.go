package spt

import (
	"fmt"

	"github.com/marselester/alg/digraph/weighted"
)

func ExampleNewDijkstra() {
	g := weighted.NewAdjacencyList(8)
	g.Add(&weighted.Edge{4, 5, 0.35})
	g.Add(&weighted.Edge{5, 4, 0.35})
	g.Add(&weighted.Edge{4, 7, 0.37})
	g.Add(&weighted.Edge{5, 7, 0.28})
	g.Add(&weighted.Edge{7, 5, 0.28})
	g.Add(&weighted.Edge{5, 1, 0.32})
	g.Add(&weighted.Edge{0, 4, 0.38})
	g.Add(&weighted.Edge{0, 2, 0.26})
	g.Add(&weighted.Edge{7, 3, 0.39})
	g.Add(&weighted.Edge{1, 3, 0.29})
	g.Add(&weighted.Edge{2, 7, 0.34})
	g.Add(&weighted.Edge{6, 2, 0.4})
	g.Add(&weighted.Edge{3, 6, 0.52})
	g.Add(&weighted.Edge{6, 0, 0.58})
	g.Add(&weighted.Edge{6, 4, 0.93})

	d := NewDijkstra(g, 0)
	for v := 0; v < g.VertexCount(); v++ {
		fmt.Printf("0 to %d: %v\n", v, d.PathTo(v))
	}
	// Output:
	// 0 to 0: []
	// 0 to 1: [0->4 0.38 4->5 0.35 5->1 0.32]
	// 0 to 2: [0->2 0.26]
	// 0 to 3: [0->2 0.26 2->7 0.34 7->3 0.39]
	// 0 to 4: [0->4 0.38]
	// 0 to 5: [0->4 0.38 4->5 0.35]
	// 0 to 6: [0->2 0.26 2->7 0.34 7->3 0.39 3->6 0.52]
	// 0 to 7: [0->2 0.26 2->7 0.34]
}

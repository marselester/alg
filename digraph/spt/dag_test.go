package spt

import "fmt"

func ExampleNewAcyclic() {
	g := NewAdjacencyList(8)
	g.Add(&Edge{5, 4, 0.35})
	g.Add(&Edge{4, 7, 0.37})
	g.Add(&Edge{5, 7, 0.28})
	g.Add(&Edge{5, 1, 0.32})
	g.Add(&Edge{4, 0, 0.38})
	g.Add(&Edge{0, 2, 0.26})
	g.Add(&Edge{3, 7, 0.39})
	g.Add(&Edge{1, 3, 0.29})
	g.Add(&Edge{7, 2, 0.34})
	g.Add(&Edge{6, 2, 0.4})
	g.Add(&Edge{3, 6, 0.52})
	g.Add(&Edge{6, 0, 0.58})
	g.Add(&Edge{6, 4, 0.93})

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

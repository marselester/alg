package mst

import (
	"fmt"
)

func ExampleNewLazyPrim() {
	g := NewAdjacencyList(8)
	g.Add(&Edge{V: 4, W: 5, Weight: 0.35})
	g.Add(&Edge{V: 4, W: 7, Weight: 0.37})
	g.Add(&Edge{V: 5, W: 7, Weight: 0.28})
	g.Add(&Edge{V: 0, W: 7, Weight: 0.16})
	g.Add(&Edge{V: 1, W: 5, Weight: 0.32})
	g.Add(&Edge{V: 0, W: 4, Weight: 0.38})
	g.Add(&Edge{V: 2, W: 3, Weight: 0.17})
	g.Add(&Edge{V: 1, W: 7, Weight: 0.19})
	g.Add(&Edge{V: 0, W: 2, Weight: 0.26})
	g.Add(&Edge{V: 1, W: 2, Weight: 0.36})
	g.Add(&Edge{V: 1, W: 3, Weight: 0.29})
	g.Add(&Edge{V: 2, W: 7, Weight: 0.34})
	g.Add(&Edge{V: 6, W: 2, Weight: 0.40})
	g.Add(&Edge{V: 3, W: 6, Weight: 0.52})
	g.Add(&Edge{V: 6, W: 0, Weight: 0.58})
	g.Add(&Edge{V: 6, W: 4, Weight: 0.93})

	mst := NewLazyPrim(g)
	for _, e := range mst.Edges() {
		fmt.Println(e)
	}
	fmt.Printf("%.5f", mst.Weight())
	// Output:
	// 0-7 0.16000
	// 1-7 0.19000
	// 0-2 0.26000
	// 2-3 0.17000
	// 5-7 0.28000
	// 4-5 0.35000
	// 6-2 0.40000
	// 1.81000
}

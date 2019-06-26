package digraph_test

import (
	"fmt"

	"github.com/marselester/alg/digraph"
)

func ExampleAdjacencyList_String() {
	g := digraph.NewAdjacencyList(13)
	g.Add(4, 2)
	g.Add(2, 3)
	g.Add(3, 2)
	g.Add(6, 0)
	g.Add(0, 1)
	g.Add(2, 0)
	g.Add(11, 12)
	g.Add(12, 9)
	g.Add(9, 10)
	g.Add(9, 11)
	g.Add(7, 9)
	g.Add(10, 12)
	g.Add(11, 4)
	g.Add(4, 3)
	g.Add(3, 5)
	g.Add(6, 8)
	g.Add(8, 6)
	g.Add(5, 4)
	g.Add(0, 5)
	g.Add(6, 4)
	g.Add(6, 9)
	g.Add(7, 6)

	fmt.Print(g.String())
	// Output:
	// 0: 5 1
	// 1:
	// 2: 0 3
	// 3: 5 2
	// 4: 3 2
	// 5: 4
	// 6: 9 4 8 0
	// 7: 6 9
	// 8: 6
	// 9: 11 10
	// 10: 12
	// 11: 4 12
	// 12: 9
}

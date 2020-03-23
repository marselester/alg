package digraph

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	g := NewAdjacencyList(13)
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

	if !HasCycle(g) {
		t.Errorf("HasCycle() = false, want true")
	}
}

func ExampleCycle() {
	g := NewAdjacencyList(13)
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

	// The expected cycle is 3->5->4->3.
	fmt.Println(Cycle(g))
	// Output:
	// [3 5 4 3]
}

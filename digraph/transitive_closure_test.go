package digraph

import (
	"fmt"
	"testing"
)

func ExampleTransitiveClosure() {
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

	tc := NewTransitiveClosure(g)
	fmt.Println(tc.Reachable(6, 12))
	fmt.Println(tc.Reachable(12, 6))
	// Output:
	// true
	// false
}

func TestTransitiveClosure(t *testing.T) {
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

	tc := NewTransitiveClosure(g)
	// want is expected reachability matrix.
	want := [][]int{
		{0, 1, 2, 3, 4, 5},
		{1},
		{0, 1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4, 5},
		{0, 1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		{0, 1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12},
		{0, 1, 2, 3, 4, 5, 9, 10, 11, 12},
		{0, 1, 2, 3, 4, 5, 9, 10, 11, 12},
		{0, 1, 2, 3, 4, 5, 9, 10, 11, 12},
		{0, 1, 2, 3, 4, 5, 9, 10, 11, 12},
	}
	for i := 0; i < len(tc.matrix); i++ {
		if !equal(tc.matrix[i], want[i]) {
			t.Fatalf("TransitiveClosure() from %d to %v vertices; want %v", i, tc.matrix[i], want[i])
		}
	}
}

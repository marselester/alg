package graph

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	g := NewAdjacencyList(13)
	g.Add(0, 1)
	g.Add(1, 2)
	g.Add(2, 3)
	g.Add(3, 4)

	if HasCycle(g) {
		t.Errorf("HasCycle() = true, want false")
	}

	g.Add(4, 0)
	if !HasCycle(g) {
		t.Errorf("HasCycle() = false, want true")
	}
}

func ExampleHasCycle() {
	g := NewAdjacencyList(5)
	g.Add(1, 0)
	g.Add(0, 2)
	g.Add(0, 3)
	g.Add(3, 4)

	fmt.Println(HasCycle(g))
	g.Add(2, 1)
	fmt.Println(HasCycle(g))
	// Output:
	// false
	// true
}

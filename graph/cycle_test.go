package graph

import "testing"

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

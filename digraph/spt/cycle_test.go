package spt

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	g := NewAdjacencyList(13)
	g.Add(&Edge{4, 2, 0})
	g.Add(&Edge{2, 3, 0})
	g.Add(&Edge{3, 2, 0})
	g.Add(&Edge{6, 0, 0})
	g.Add(&Edge{0, 1, 0})
	g.Add(&Edge{2, 0, 0})
	g.Add(&Edge{11, 12, 0})
	g.Add(&Edge{12, 9, 0})
	g.Add(&Edge{9, 10, 0})
	g.Add(&Edge{9, 11, 0})
	g.Add(&Edge{7, 9, 0})
	g.Add(&Edge{10, 12, 0})
	g.Add(&Edge{11, 4, 0})
	g.Add(&Edge{4, 3, 0})
	g.Add(&Edge{3, 5, 0})
	g.Add(&Edge{6, 8, 0})
	g.Add(&Edge{8, 6, 0})
	g.Add(&Edge{5, 4, 0})
	g.Add(&Edge{0, 5, 0})
	g.Add(&Edge{6, 4, 0})
	g.Add(&Edge{6, 9, 0})
	g.Add(&Edge{7, 6, 0})

	if !hasCycle(g) {
		t.Errorf("hasCycle() = false, want true")
	}
}

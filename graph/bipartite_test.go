package graph

import "testing"

func TestIsBipartite(t *testing.T) {
	g := NewAdjacencyList(13)
	// Red vertices: 0, 3, 4, 7, 10, 11.
	g.Add(0, 1)
	g.Add(0, 2)
	g.Add(0, 5)
	g.Add(0, 6)
	g.Add(1, 3)
	g.Add(3, 5)
	g.Add(4, 5)
	g.Add(4, 6)
	g.Add(6, 7)
	g.Add(7, 8)
	g.Add(8, 10)
	g.Add(9, 10)
	g.Add(9, 11)
	g.Add(10, 12)
	g.Add(11, 12)

	if !IsBipartite(g) {
		t.Errorf("IsBipartite() = false, want true")
	}
}

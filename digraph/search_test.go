package digraph

import (
	"testing"
)

func TestDepthFirstSearch(t *testing.T) {
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

	tests := []struct {
		g         *AdjacencyList
		source    []int
		reachable []int
	}{
		{
			g:         g,
			source:    []int{1},
			reachable: []int{1},
		},
		{
			g:         g,
			source:    []int{2},
			reachable: []int{0, 1, 2, 3, 4, 5},
		},
		{
			g:         g,
			source:    []int{1, 2, 6},
			reachable: []int{0, 1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12},
		},
	}
	for _, tc := range tests {
		got := DepthFirstSearch(g, tc.source...)
		if !equal(got, tc.reachable) {
			t.Errorf("DepthFirstSearch(%v) = %v, want %v", tc.source, got, tc.reachable)
		}
	}
}

func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

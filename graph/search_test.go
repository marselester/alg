package graph

import (
	"testing"
)

func equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestDepthFirstSearch(t *testing.T) {
	g := NewAdjacencyList(13)
	g.Add(0, 5)
	g.Add(4, 3)
	g.Add(0, 1)
	g.Add(9, 12)
	g.Add(6, 4)
	g.Add(5, 4)
	g.Add(0, 2)
	g.Add(11, 12)
	g.Add(9, 10)
	g.Add(0, 6)
	g.Add(7, 8)
	g.Add(9, 11)
	g.Add(5, 3)

	connected := DepthFirstSearch(g, 0)
	want := []int{0, 1, 2, 3, 4, 5, 6}
	if !equal(connected, want) {
		t.Errorf("DepthFirstSearch(0) = %v, want %v", connected, want)
	}

	connected = DepthFirstSearch(g, 9)
	want = []int{9, 10, 11, 12}
	if !equal(connected, want) {
		t.Errorf("DepthFirstSearch(9) = %v, want %v", connected, want)
	}
}

func TestDepthFirstPath(t *testing.T) {
	g := NewAdjacencyList(6)
	g.Add(0, 5)
	g.Add(2, 4)
	g.Add(2, 3)
	g.Add(1, 2)
	g.Add(0, 1)
	g.Add(3, 4)
	g.Add(3, 5)
	g.Add(0, 2)

	path := NewDepthFirstPath(g, 0).To(5)
	want := []int{0, 2, 3, 5}
	if !equal(path, want) {
		t.Errorf("DepthFirstPath() from 0 to 5 = %v, want %v", path, want)
	}

	g = NewAdjacencyList(4)
	g.Add(0, 1)
	g.Add(2, 3)

	path = NewDepthFirstPath(g, 0).To(2)
	want = nil
	if !equal(path, want) {
		t.Errorf("DepthFirstPath() from 0 to 2 = %v, want %v", path, want)
	}

	path = NewDepthFirstPath(g, 2).To(3)
	want = []int{2, 3}
	if !equal(path, want) {
		t.Errorf("DepthFirstPath() from 2 to 3 = %v, want %v", path, want)
	}
}

func TestBreadthFirstPath(t *testing.T) {
	g := NewAdjacencyList(6)
	g.Add(0, 5)
	g.Add(2, 4)
	g.Add(2, 3)
	g.Add(1, 2)
	g.Add(0, 1)
	g.Add(3, 4)
	g.Add(3, 5)
	g.Add(0, 2)

	path := NewBreadthFirstPath(g, 0).To(4)
	want := []int{0, 2, 4}
	if !equal(path, want) {
		t.Errorf("BreadthFirstPath() from 0 to 4 = %v, want %v", path, want)
	}
}

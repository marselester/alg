package spt

import (
	"fmt"
	"testing"
)

func TestRelax(t *testing.T) {
	e := &Edge{V: 0, W: 1}
	// Length of the shortest path from source to v is distTo[e.V] = 3.1,
	// length of the shortest path from source to w is distTo[e.W] = 3.3.
	distTo := make([]float64, 2, 2)
	distTo[e.V] = 3.1
	distTo[e.W] = 3.3

	// When weight is 0.1, edge e leads to a shorter path to w, because 3.3 > 3.1 + 0.1.
	e.Weight = 0.1
	edgeTo := make([]*Edge, 2, 2)
	if ok := relax(e, distTo, edgeTo); !ok {
		t.Errorf("relax() = false when dist to v + edge weight < dist to w")
	}
	if edgeTo[e.V] != nil || edgeTo[e.W] != e {
		t.Errorf("relax() edgeTo was not updated: %v", edgeTo)
	}

	// When e.Weight is 1.3, edge e is ignored, because 3.3 < 3.1 + 1.3.
	e.Weight = 1.3
	edgeTo = make([]*Edge, 2, 2)
	if ok := relax(e, distTo, edgeTo); ok {
		t.Errorf("relax() = true when dist to v + edge weight > dist to w")
	}
	if edgeTo[e.V] != nil || edgeTo[e.W] != nil {
		t.Errorf("relax() edgeTo was updated: %v", edgeTo)
	}
}

func TestPathTo(t *testing.T) {
	edgeTo := []*Edge{
		nil,
		{5, 1, 0.32},
		{0, 2, 0.26},
		{7, 3, 0.39},
		{0, 4, 0.38},
		{4, 5, 0.35},
		{3, 6, 0.52},
		{2, 7, 0.34},
	}
	want := []*Edge{
		{0, 2, 0.26},
		{2, 7, 0.34},
		{7, 3, 0.39},
		{3, 6, 0.52},
	}

	got := pathTo(6, edgeTo)
	if fmt.Sprint(want) != fmt.Sprint(got) {
		t.Errorf("pathTo(6, %v) = %v want %v", edgeTo, got, want)
	}
}

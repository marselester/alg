package graph

import "testing"

func TestConnectedComponent(t *testing.T) {
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

	cc := NewConnectedComponent(g)
	var compid int
	components := make([][]int, cc.Count())
	for v := 0; v < 13; v++ {
		compid = cc.ID(v)
		components[compid] = append(components[compid], v)
	}

	want := [][]int{
		{0, 1, 2, 3, 4, 5, 6},
		{7, 8},
		{9, 10, 11, 12},
	}
	for compid = 0; compid < len(components); compid++ {
		if !equal(components[compid], want[compid]) {
			t.Errorf("ConnectedComponent(%d) = %v, want %v", compid, components[compid], want[compid])
		}
	}
}

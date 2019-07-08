package digraph

import "testing"

func TestStrongConnectedComponent(t *testing.T) {
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

	scc := NewStrongConnectedComponent(g)
	if scc.Count() != 5 {
		t.Fatalf("StrongComponent() found %d components, want 5", scc.Count())
	}

	var compid int
	components := make([][]int, scc.Count())
	for v := 0; v < g.VertexCount(); v++ {
		compid = scc.ID(v)
		components[compid] = append(components[compid], v)
	}

	want := [][]int{
		{1},
		{0, 2, 3, 4, 5},
		{9, 10, 11, 12},
		{6, 8},
		{7},
	}
	for compid = 0; compid < len(components); compid++ {
		if !equal(components[compid], want[compid]) {
			t.Errorf("StrongComponent(%d) = %v, want %v", compid, components[compid], want[compid])
		}
	}
}

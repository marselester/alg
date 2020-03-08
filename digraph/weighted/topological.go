package weighted

// TopologicalSort computes a topological order for the vertices of a DAG (directed acyclic graph).
func TopologicalSort(g *AdjacencyList) []int {
	if HasCycle(g) {
		return nil
	}
	return reversePostorder(g)
}

func reversePostorder(g *AdjacencyList) []int {
	t := topological{
		g:           g,
		marked:      make([]bool, g.VertexCount()),
		reversePost: make([]int, g.VertexCount()),
	}
	for v := 0; v < g.VertexCount(); v++ {
		if !t.marked[v] {
			t.dfs(v)
		}
	}
	return t.reversePost[:g.VertexCount()]
}

type topological struct {
	g           *AdjacencyList
	marked      []bool
	reversePost []int
}

func (t *topological) dfs(source int) {
	t.marked[source] = true
	for _, e := range t.g.Adjacent(source) {
		if !t.marked[e.W] {
			t.dfs(e.W)
		}
	}

	last := len(t.reversePost) - 1
	t.reversePost[last] = source
	t.reversePost = t.reversePost[:last]
}

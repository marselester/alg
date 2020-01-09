package spt

func hasCycle(g *AdjacencyList) bool {
	c := cycle{
		g:       g,
		onStack: make([]bool, g.VertexCount()),
		marked:  make([]bool, g.VertexCount()),
	}
	// It finds an unmarked vertex and calls the recursive DFS to mark and identify
	// all vertices connected to it, continuing until all vertices have been marked and identified.
	for v := 0; v < g.VertexCount(); v++ {
		if !c.marked[v] {
			c.dfs(v)
		}
	}
	return c.exists
}

type cycle struct {
	g *AdjacencyList
	// onStack is used to track the vertices for which recursive dfs call hasn't completed.
	onStack []bool
	// marked is an array of visited vertices.
	marked []bool
	// exists indicates whether a graph has a directed cycle.
	// If no cycle exists, a digraph is a directed acyclic graph (DAG).
	exists bool
}

func (c *cycle) dfs(source int) {
	c.onStack[source] = true

	c.marked[source] = true
	for _, e := range c.g.Adjacent(source) {
		switch {
		case c.exists:
			return
		case !c.marked[e.W]:
			c.dfs(e.W)
		case c.onStack[e.W]:
			c.exists = true
		}
	}

	c.onStack[source] = false
}

package weighted

// HasCycle reports whether a given graph has a directed cycle.
func HasCycle(g *AdjacencyList) bool {
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

// Cycle returns a first cycle found in the directed graph. When cycle is detected,
// it finds the vertices from last known source vertex back to itself.
func Cycle(g *AdjacencyList) []*Edge {
	c := cyclepath{
		g:       g,
		onStack: make([]bool, g.VertexCount()),
		marked:  make([]bool, g.VertexCount()),
		edgeTo:  make([]*Edge, g.VertexCount()),
	}
	for s := range g.a {
		if !c.marked[s] {
			c.dfs(s)
		}
	}

	for i, j := 0, len(c.cycle)-1; i < j; i, j = i+1, j-1 {
		c.cycle[i], c.cycle[j] = c.cycle[j], c.cycle[i]
	}
	return c.cycle
}

type cyclepath struct {
	g *AdjacencyList
	// onStack is used to track the vertices for which recursive dfs call hasn't completed.
	onStack []bool
	// marked is an array of visited vertices.
	marked []bool
	// edgeTo holds a last vertex on known path to a vertex.
	edgeTo []*Edge
	// cycle is the vertices that form a cycle.
	cycle []*Edge
}

func (c *cyclepath) dfs(source int) {
	c.onStack[source] = true

	c.marked[source] = true
	for _, e := range c.g.Adjacent(source) {
		switch {
		// Cycle was already found.
		case c.cycle != nil:
			return
		// Found the unvisited vertex, need to identify all vertices connected to it
		// and remember how to get there (path to the unmarked vertex is from the source).
		case !c.marked[e.W]:
			c.edgeTo[e.W] = e
			c.dfs(e.W)
		case c.onStack[e.W]:
			var x *Edge
			for x = e; x.V != e.W; x = c.edgeTo[x.V] {
				c.cycle = append(c.cycle, x)
			}
			c.cycle = append(c.cycle, x)
		}
	}

	c.onStack[source] = false
}

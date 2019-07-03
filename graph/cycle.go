package graph

// HasCycle reports whether a given graph has a cycle (path with at least one edge
// whose first and last vertices are the same). It assumes no self-loops or parallel edges.
func HasCycle(g *AdjacencyList) bool {
	c := cycle{
		g:      g,
		marked: make([]bool, len(g.a)),
	}
	// It finds an unmarked vertex and calls the recursive DFS to mark and identify
	// all vertices connected to it, continuing until all vertices have been marked and identified.
	for s := range g.a {
		if !c.marked[s] {
			c.dfs(s, -1)
		}
	}
	return c.exists
}

type cycle struct {
	g *AdjacencyList
	// marked is an array of visited vertices.
	marked []bool
	// exists indicates whether a graph has a cycle.
	exists bool
}

/*
dfs starts visiting vertices from source vertex and its parent vertex (ex source).
For example, if path is 0->5->4->3->5(check), then dfs call stack is:

	0,-1 -> 5,0 -> 4,5 -> source=3,parent=4 (v=5 is already marked)

There is a cycle, because vertex 5 is already visited,
and adjacent visited vector is not parent of v.
*/
func (c *cycle) dfs(source, parent int) {
	c.marked[source] = true
	for n := c.g.a[source]; n != nil; n = n.next {
		if !c.marked[n.v] {
			c.dfs(n.v, source)
		} else
		// Cycle: adjacent vertex v is visited and v is not parent of source.
		if n.v != parent {
			c.exists = true
		}
	}
}

package digraph

/*
HasCycle reports whether a given graph has a directed cycle.

In a scheduling problem certain jobs must be performed before certain others (precedence constraint)
which is a topological sort problem. No solution exists if there is a directed cycle, for example,
those three constraints cannot all be satisfied:

	job X must be completed before job Y
	job Y before job Z
	job Z before job X

A depth-first search solution is based on the fact that the recursive function call stack represents
the "current" directed path under consideration.
If we ever find a directed edge to a vertex that is on that stack, we have found a cycle,
since the stack is evidence of a directed path.
For example, in 0->5->4->3->5(check) call stack, the last "5" completes the cycle (3->5->4->3).
*/
func HasCycle(g *AdjacencyList) bool {
	c := cycle{
		g:       g,
		onStack: make([]bool, g.VertexCount()),
		marked:  make([]bool, g.VertexCount()),
	}
	// It finds an unmarked vertex and calls the recursive DFS to mark and identify
	// all vertices connected to it, continuing until all vertices have been marked and identified.
	for s := range g.a {
		if !c.marked[s] {
			c.dfs(s)
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
	for n := c.g.a[source]; n != nil; n = n.next {
		switch {
		case c.exists:
			return
		case !c.marked[n.v]:
			c.dfs(n.v)
		case c.onStack[n.v]:
			c.exists = true
		}
	}

	c.onStack[source] = false
}

// Cycle returns a first cycle found in the directed graph. When cycle is detected,
// it finds the vertices from last known source vertex back to itself.
func Cycle(g *AdjacencyList) []int {
	c := cyclepath{
		g:       g,
		onStack: make([]bool, g.VertexCount()),
		marked:  make([]bool, g.VertexCount()),
		edgeTo:  make([]int, g.VertexCount()),
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
	edgeTo []int
	// cycle is the vertices that form a cycle.
	cycle []int
}

func (c *cyclepath) dfs(source int) {
	c.onStack[source] = true

	c.marked[source] = true
	for n := c.g.a[source]; n != nil; n = n.next {
		switch {
		// Cycle was already found.
		case c.cycle != nil:
			return
		// Found the unvisited vertex, need to identify all vertices connected to it
		// and remember how to get there (path to the unmarked vertex is from the source).
		case !c.marked[n.v]:
			c.edgeTo[n.v] = source
			c.dfs(n.v)
		// Cycle has just been detected, need to backtrack the path.
		// For example, in 0->5->4->3->5(check) call stack, the last "5" completes the cycle (3->5->4->3),
		// where source=3 and n.v=5.
		case c.onStack[n.v]:
			for x := source; x != n.v; x = c.edgeTo[x] {
				c.cycle = append(c.cycle, x)
			}
			c.cycle = append(c.cycle, n.v, source)
		}
	}

	c.onStack[source] = false
}

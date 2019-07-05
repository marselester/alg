package digraph

/*
TopologicalSort computes a topological order for the vertices of a DAG (directed acyclic graph).
A digraph has a topological order if and only if it's a DAG.
The idea is that DFS visits each vertex exactly once.
If we save the vertex given as argument to the recursive dfs() in a data structure,
in order by where we do the save (before or after the recursive calls):

	Preorder: Put the vertex on a queue before the recursive calls.
	Postorder: Put the vertex on a queue after the recursive calls.
	Reverse postorder: Put the vertex on a stack after the recursive calls.

Reverse postorder in a DAG is a topological sort. With DFS, we can topologically sort a DAG in time
proportional to E+V. It uses one depth-first search to ensure the graph has no directed cycles,
and another to do the reverse postorder ordering.
*/
func TopologicalSort(g *AdjacencyList) []int {
	if HasCycle(g) {
		return nil
	}

	t := topological{
		g:           g,
		marked:      make([]bool, g.VertexCount()),
		reversePost: make([]int, g.VertexCount()),
	}
	for s := range g.a {
		if !t.marked[s] {
			t.dfs(s)
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
	for n := t.g.a[source]; n != nil; n = n.next {
		if !t.marked[n.v] {
			t.dfs(n.v)
		}
	}

	last := len(t.reversePost) - 1
	t.reversePost[last] = source
	t.reversePost = t.reversePost[:last]
}

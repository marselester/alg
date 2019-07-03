package digraph

/*
DepthFirstSearch is a method to find vertices that are reachable from a given vertex
to support queries "is there a directed path from source to a given target vertex v?".

Multiple-source reachability: given a digraph and a set of source vertices,
support queries of the form "is there a directed path from some vertex in the set
to a given target vertex v". For example, it's applicable to mark-and-sweep garbage collection:

	- vertex represents a variable, edge is a reference to variable
	- certain variables are known to be directly accessible
	- any variables that are not reachable from that set can be returned to available memory
	  ("not reachable" means there is no "->" from the set to a variable)

Depth-first search marks all the vertices reachable from a given set of sources in time
proportional to the sum of thei outdegrees of the vertices marked.
The outdegree of a vertex v is a number of edges leaving it.
*/
func DepthFirstSearch(g *AdjacencyList, source ...int) []int {
	marked := make([]bool, len(g.a))
	for _, v := range source {
		dfs(g, v, marked)
	}

	var reachable []int
	for v := range marked {
		if marked[v] {
			reachable = append(reachable, v)
		}
	}
	return reachable
}

// dfs visits all vertices connected to the source vertex. To visit a vertex,
// mark it as visited and visit (recursively) all the vertices that are adjacent to it and
// that have not been marked yet.
func dfs(g *AdjacencyList, source int, marked []bool) {
	marked[source] = true
	for n := g.a[source]; n != nil; n = n.next {
		if !marked[n.v] {
			dfs(g, n.v, marked)
		}
	}
}

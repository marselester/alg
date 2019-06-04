package graph

/*
IsBipartite returns true if a graph whose vertices we can divide into two sets
such that all edges connect a vertex in one set (colored red) with a vertex in the other set
(colored black). In other words, can the vertices of a given graph be assigned one of two colors
in such a way that no edge connects vertices of the same color?

For example, IMDB can be defined as a graph with movies and performers as vertices and
each line defining the adjacency list of edges connecting each movie to its performers.
The graph is bipartite — there are no edges connecting performers to performers or movies to movies.
*/
func IsBipartite(g *AdjacencyList) bool {
	b := bipartite{
		g:            g,
		marked:       make([]bool, len(g.a)),
		color:        make([]bool, len(g.a)),
		is2colorable: true,
	}
	// It finds an unmarked vertex and calls the recursive DFS to mark and identify
	// all vertices connected to it, continuing until all vertices have been marked and identified.
	for s := range g.a {
		if !b.marked[s] {
			b.dfs(s)
		}
	}
	return b.is2colorable
}

type bipartite struct {
	g *AdjacencyList
	// marked is an array of visited vertices.
	marked []bool
	// color represents vertices' color.
	color        []bool
	is2colorable bool
}

func (b *bipartite) dfs(source int) {
	b.marked[source] = true
	for n := b.g.a[source]; n != nil; n = n.next {
		if !b.marked[n.v] {
			b.color[n.v] = !b.color[source]
			b.dfs(n.v)
		} else if b.color[n.v] == b.color[source] {
			b.is2colorable = false
		}
	}
}

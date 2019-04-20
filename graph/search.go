package graph

/*
DepthFirstSearch is a method to find the vertices connected to the source.
It is similar to searching through the maze by using a string that
guarantees you can always find a way out and the passage marks guarantee that
you avoid exploring any passage or intersection twice.
To search a graph, a recursive method is used that visits all of graph's vertices and edges.

Depth-first search marks all the vertices connected to a given source in time proportional
to the sum of their degrees. A degree of vertex v is a count of its adjacent vertices.
*/
func DepthFirstSearch(g *AdjacencyList, source int) []int {
	// marked is an array of visited vertices connected to the source vertex.
	marked := make([]bool, len(g.a))
	search(g, source, marked)

	var connected []int
	for v := range marked {
		if marked[v] {
			connected = append(connected, v)
		}
	}
	return connected
}

// search visits all vertices connected to the source vertex. To visit a vertex
// mark it as visited and visit (recursively) all the vertices that are adjacent to it and
// that have not been marked yet.
func search(g *AdjacencyList, source int, marked []bool) {
	marked[source] = true
	for n := g.a[source]; n != nil; n = n.next {
		if !marked[n.v] {
			search(g, n.v, marked)
		}
	}
}

// NewDepthFirstPath creates DepthFirstPath to look up paths connected to source vertex.
func NewDepthFirstPath(g *AdjacencyList, source int) *DepthFirstPath {
	dfs := DepthFirstPath{
		g:      g,
		source: source,
		marked: make([]bool, len(g.a)),
		edgeTo: make([]int, len(g.a)),
	}
	dfs.searchPaths(source)
	return &dfs
}

/*
DepthFirstPath uses depth-first search to find paths to all the vertices in a graph
that are connected to a given start vertex s. To save known paths to each vertex,
it maintains a vertex-indexed array.
*/
type DepthFirstPath struct {
	g *AdjacencyList
	// source is a vertex to which connections are discovered.
	source int
	// marked is an array of visited vertices connected to the source vertex.
	marked []bool
	// edgeTo is a parent-link representation of a tree rooted at source vertex
	// that contains all the vertices connected to source.
	edgeTo []int
}

// searchPaths visits all vertices connected to the source and saves known paths to each vertex.
// To visit a vertex mark it as visited and visit (recursively) all the vertices
// that are adjacent to it and that have not been marked yet.
func (dfs *DepthFirstPath) searchPaths(source int) {
	dfs.marked[source] = true
	for n := dfs.g.a[source]; n != nil; n = n.next {
		if !dfs.marked[n.v] {
			// edgeTo is a vertex-indexed array such that edgeTo[b] = a means
			// a-b edge was used to access b vertex for the first time.
			dfs.edgeTo[n.v] = source
			dfs.searchPaths(n.v)
		}
	}
}

// To searches a path from source to the connected vertex v.
func (dfs *DepthFirstPath) To(v int) []int {
	// No such path exists.
	if !dfs.marked[v] {
		return nil
	}

	// array index is a destination vertex, value is a start vertex.
	// Start from destination vertex v and trace back the path to the source.
	var path []int
	for ; v != dfs.source; v = dfs.edgeTo[v] {
		path = append(path, v)
	}
	path = append(path, dfs.source)

	// Reverse the path so it starts from source vertex and ends at v vertex.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// NewBreadthFirstPath creates BreadthFirstPath to look up shortest paths from the source vertex.
func NewBreadthFirstPath(g *AdjacencyList, source int) *BreadthFirstPath {
	bfs := BreadthFirstPath{
		g:      g,
		source: source,
		marked: make([]bool, len(g.a)),
		edgeTo: make([]int, len(g.a)),
	}
	bfs.searchPaths(source)
	return &bfs
}

/*
BreadthFirstPath uses breadth-first search to find paths in a graph
with the fewest number of edges from the source vertex.
The algorithm is based on maintaining a queue of all vertices that have been marked
but whose adjacency lists have not been checked.
It puts the source vertex onto the queue, then does the following steps until the queue is empty:

	- remove the next vertex v from the queue,
	- put onto the queue all unmarked vertices that are adjacent to v and mark them.

BFS takes time proportional to number of edges + number of vertices in the worst case.
*/
type BreadthFirstPath struct {
	g *AdjacencyList
	// source is a vertex to which connections are discovered.
	source int
	// marked is an array of vertices to which a shortest path is known.
	marked []bool
	// edgeTo represents a last vertex on known path to this vertex.
	edgeTo []int
}

// searchPaths looks up shortest paths from the source vertex.
func (bfs *BreadthFirstPath) searchPaths(source int) {
	bfs.marked[source] = true
	q := queue{}
	q.Enqueue(source)
	for q.Size() != 0 {
		v := q.Dequeue()
		for n := bfs.g.a[v]; n != nil; n = n.next {
			if !bfs.marked[n.v] {
				bfs.edgeTo[n.v] = v
				q.Enqueue(n.v)
				bfs.marked[n.v] = true
			}
		}
	}
}

// To searches a shortest path from source to the connected vertex v.
func (bfs *BreadthFirstPath) To(v int) []int {
	// No such path exists.
	if !bfs.marked[v] {
		return nil
	}

	// array index is a destination vertex, value is a start vertex.
	// Start from destination vertex v and trace back the path to the source.
	var path []int
	for ; v != bfs.source; v = bfs.edgeTo[v] {
		path = append(path, v)
	}
	path = append(path, bfs.source)

	// Reverse the path so it starts from source vertex and ends at v vertex.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// queue represents a queue of vertices.
type queue struct {
	vertices []int
}

// Enqueue adds a vertex to the end of the queue.
func (q *queue) Enqueue(v int) {
	q.vertices = append(q.vertices, v)
}

// Dequeue returns a vertex from the beginning of the queue.
func (q *queue) Dequeue() int {
	if len(q.vertices) == 0 {
		return 0
	}
	v := q.vertices[0]
	q.vertices = q.vertices[1:]
	return v
}

// Size returns the number of vertices in the queue.
func (q *queue) Size() int {
	return len(q.vertices)
}

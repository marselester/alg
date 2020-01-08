package spt

// relax performs edge-relaxation operation (relaxing the tension on the rubber band along a shorter path)
// and reports if the edge was eligible.
//
// To relax an edge v->w means to test whether the best known way from source to w is to go from source to v,
// then take the edge from v to w.
// The best known distance to w through v is the sum of distTo[v] and edge weight:
//
//     if the sum is smaller than distTo[w], then edge v->w is eligible
//     otherwise v->w is ineligible (ignored)
//
// For example, length of the shortest path from source to v is distTo[e.V] = 3.1,
// length of the shortest path from source to w is distTo[e.W] = 3.3.
// If e.Weight is 1.3, then edge e is ignored, because 3.3 < 3.1 + 1.3.
// If weight is 0.1, then edge e leads to a shorter path to w, because 3.3 > 3.1 + 0.1.
func relax(e *Edge, distTo []float64, edgeTo []*Edge) bool {
	distance := distTo[e.V] + e.Weight
	isEligible := distTo[e.W] > distance
	if isEligible {
		distTo[e.W] = distance
		edgeTo[e.W] = e
	}
	return isEligible
}

// pathTo returns a path from source vertex to v.
// Note, a caller should make sure the path exists.
func pathTo(v int, edgeTo []*Edge) []*Edge {
	// path's array index is a destination vertex, value is a start vertex.
	// Start from destination vertex v and trace back the path to the source.
	var path []*Edge
	for e := edgeTo[v]; e != nil; e = edgeTo[e.V] {
		path = append(path, e)
	}

	// Reverse the path so it starts from source vertex and ends at v vertex.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

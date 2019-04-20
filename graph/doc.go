/*
Package graph provides data structures to manipulate a set of vertices and
a collection of edges that each connect a pair of vertices.

Depth-first search (DFS) and breadth-first search (BFS) start from source vertex
and perform the following steps until the data structure is empty:

	- take the next unmarked vertex v from the data structure and mark it,
	- put onto the data structure all unmarked vertices that are adjacent to v.

The algorithms differ only in the rule used to take the next vertex:
most recently added for DFS (stack), least recently added for BFS (queue).

DFS explores the graph by looking for new vertices far away from the start point,
taking closer vertices only when dead ends are encountered.

BFS completely covers the area close to the starting point, moving farther away
only when everything nearby has been examined.
*/
package graph

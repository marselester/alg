/*
Package mst computes minimum spanning tree for undirected edge-weighted graph.
Weight represents a cost (or time) that it takes to propagate through the edge
(airline map where weights are distances or fares).
Minimizing the cost is naturally of interest in such situations.

A spanning tree of a graph is a connected subgraph with no cycles that includes all the vertices.
A min spanning tree (MST) of a graph is a spanning tree whose weights (sum)
is no larger than weight of any other spanning tree.

A cut of a graph is a partition of its vertices into two nonempty disjoint sets.
A crossing edge of a cut is an edge that connects a vertex in one set with a vertex in the other.
Min-weight crossing edge must be in the MST.

The cut property is the basis for Prim's and Kruskal's algorithms for the MST problem.
They are special cases of a general paradigm known as greedy algorithm:
apply the cut property to accept an edge as an MST edge,
continuing until finding all of the MST edges.
*/
package mst

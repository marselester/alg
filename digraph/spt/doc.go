// Package spt computes shortest-path tree (SPT) which gives a shortest path from a source vertex s
// to every vertex reachable from s.
// The problem is formulated as finding a lowest-cost way to get from one vertex to another.
// For example, a navigation system to get directions, arbitrage problem from computational finance.
//
// Generic shortest-paths algorithm: init distTo[s] to 0 and all other distTo[] values to infinity,
// then relax any edge in digraph, continuing until no edge is eligible.
// For all vertices w reachable from s, the value of distTo[w] is the length of a shortest path from s to w
// (and edgeTo[w] is the last edge on such a path).
package spt

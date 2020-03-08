// Package lpt focuses on the problem of finding the longest path in an edge-weighted DAG
// with edge weights that may be positive or negative.
// This allows to solve parallel precedence-constrained scheduling problem:
// a set of jobs of specified duration to be completed on identical processors in the minimum time.
//
// It uses shortest-paths model (Topological sort) to solve the problem.
// Mods for shortest-paths algorithm: copy DAG with negated weights, then the shortest path in this copy
// is the longest path in the original. Simpler way to implement: switch distTo init to negative infinity,
// and change the sense of the inequality in relax().
package lpt

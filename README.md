# Algorithms

[![Documentation](https://godoc.org/github.com/marselester/alg?status.svg)](https://godoc.org/github.com/marselester/alg)
[![Go Report Card](https://goreportcard.com/badge/github.com/marselester/alg)](https://goreportcard.com/report/github.com/marselester/alg)

Go implementation of some examples from [Algorithms book](https://algs4.cs.princeton.edu).

Table of contents:

- [Sorting](#sorting)
  - [Priority queue](#priority-queue)
  - [String sorts](#string-sorts)
- [Searching](#searching)
  - [String symbol-table](#string-symbol-table)
  - [Substring search](#substring-search)
- [Dynamic connectivity](#dynamic-connectivity)
- [Graph](#graph)
  - [Minimum spanning tree](#minimum-spanning-tree)
- [Digraph](#digraph)
  - [Shortest paths](#shortest-paths)
  - [Longest paths](#longest-paths)

## Sorting

| algorithm       | running time | notes
| ---             | ---          | ---
| selection sort  | n²           |
| insertion sort  | between n and n² (depends on order of items) | stable
| shellsort       | n log n ?    |
| quicksort       | n log n (probabilistic guarantee) | extra space lg n
| 3-way quicksort | between n and n log n | extra space lg n
| mergesort       | n log n      | stable, not in place, extra space n
| heapsort        | n log n      |

A sorting method is stable if it preserves the relative order of equal keys in the array.
For example, transactions sorted by location will still preserve order by timestamp.

It's reasonable to avoid the costs of using references and sort primitive types instead.

Quicksort is the fastest general-purpose sort when space is tight:

- it has only a few instructions in its inner loop
- it does well with cache memories because it most often references data sequentially

3-way quicksort is suitable for large numbers of equal keys.

Insertion sort is an excellent method for partially sorted arrays and is also
a fine method for tiny arrays (~15 items). These properties can be leveraged in
intermediate stages of mergesort and quicksort.

Heapsort is popular when space is tight (embedded systems), but it has poor cache performance:
array entries are rarely compared with nearby array entries, so the number of cache misses is far
higher than for quicksort, mergesort, shellsort.

Mergesort is a general-purpose stable sort.

### Priority queue

Often, we accumulate items, then process the one with the largest key,
and collect more items and process the current largest key.
For example, process scheduler picks a process with the highest priority.

| data structure | insert | remove maximum | change priority
| ---            | ---    | ---            | ---
| binary heap    | log n  | log n          | log n

Examples:

- [MaxHeap](https://godoc.org/github.com/marselester/alg/sort/pqueue#MaxHeap)
  returns any largest item
- [MinHeap](https://godoc.org/github.com/marselester/alg/sort/pqueue#MinHeap)
  returns any smallest item
- [IndexMinHeap](https://godoc.org/github.com/marselester/alg/sort/pqueue#IndexMinHeap)
  allows clients to refer to items on priority queue

### String sorts

| algorithm       | running time | notes
| ---             | ---          | ---
| LSD string sort | strlen * n   | stable, not in place, extra space n
| MSD string sort | between n and strlen * n | stable, not in place, extra space n + strlen * radix
| 3-way string quicksort | between n and strlen * n * log radix | extra space strlen + log n

LSD string sort is for short fixed-length strings.

MSD string sort is for random strings.

3-way string quicksort is a general-purpose inplace sort that does well on strings with long prefix matches.

## Searching

| data structure (algorithm)      | pros | cons
| ---                             | ---  | ---
| linked list (sequential search) | best for tiny symbol tables | slow for large symbol tables
| ordered array (binary search)   | optimal search and space, order-based ops | slow insert
| binary search tree              | easy to implement, order-based ops | no guarantess, space for links
| balanced binary search tree     | optimal search and insert, order-based ops | space for links
| hash table                      | fast search/insert | hash for each type, space for links/empty

Worst and average-case costs for symbol-table implementations.

| data structure (algorithm)      | worst (search, insert) | average (search, insert)
| ---                             | ---                    | ---
| linked list (sequential search) | n                      | n/2, n
| ordered array (binary search)   | lg n, n                | lg n, n/2
| binary search tree              | n                      | 1.39 lg n
| red-black BST                   | 2 lg n                 | lg n
| separate chaining               | n                      | n/(2*m), n/m
| linear probing                  | n                      | < 1.5, < 2.5

### String symbol-table

| data structure (algorithm) | sweet spot
| ---                        | ---
| binary search tree (BST)   | randomly ordered keys
| red-black BST              | guaranteed performance
| linear probing             | built-in types, cached hash values
| R-way trie                 | short keys, small alphabets
| ternary search trie (TST)  | nonrandom keys

If space is available, R-way tries provide the fastest search (a constant number of character compares).

For large alphabets, where space may not be available for R-way tries,
TSTs (ternary search tries) are prefereable, since they use a logarithmic number of character compares,
while BSTs use a logarithmic number of key compares.

Hashing can be competitive, but cannot support ordered symbol table operations.

### Substring search

Note, `m` is a pattern length, `n` is a text length.

Brute-force search is easy to implement and works well in typical cases,
but might require time proportional to m*n.

Knuth-Morris-Pratt is guaranteed linear-time with no backup in the input,
but uses extra space.

Boyer-Moore is sublinear (by a factor of m) in typical situations,
but uses extra space.

Rabin-Karp is linear, but has a relatively long inner loop
(several arithmetic operations, as opposed to character compares in the other methods).

## Dynamic connectivity

The input is a sequence of int pairs (p, q), where each integer represents an object (computer in network)
and task is to learn if p is connected to q (connection in network).

| algorithm | union | find
| ---       | ---   | ---
| [quick-find](https://godoc.org/github.com/marselester/alg/unionfind/qfind) | n | 1
| [quick-union](https://godoc.org/github.com/marselester/alg/unionfind/qunion) | tree height | tree height
| [weighted quick-union](https://godoc.org/github.com/marselester/alg/unionfind/wqunion) | log n | log n

## Graph

Undirected graph-processing problems and solutions:

- **single-source connectivity**
  ([depth-first search](https://godoc.org/github.com/marselester/alg/graph#DepthFirstSearch))
  — find vertices connected to a given vertex
- **single-source paths**
  ([depth-first paths](https://godoc.org/github.com/marselester/alg/graph#DepthFirstPath))
  — find paths to all the vertices that are connected to a given vertex
- **single-source shortest paths**
  ([breadth-first paths](https://godoc.org/github.com/marselester/alg/graph#BreadthFirstPath))
  — find paths with the fewest number of edges from a given vertex
- **connectivity**
  ([connected components](https://godoc.org/github.com/marselester/alg/graph#ConnectedComponent))
  — find the connected components of a graph to tell whether v and w vertices are connected.
- **cycle detection**
  ([has cycle](https://godoc.org/github.com/marselester/alg/graph#HasCycle))
  — path with at least one edge whose first and last vertices are the same
- **two-colorability**
  ([is bipartite](https://godoc.org/github.com/marselester/alg/graph#IsBipartite))
  — can the vertices be assigned one of two colors in such a way
  that no edge connects vertices of the same color?
  For example, IMDB can be defined as a graph with movies and performers.
  The graph is bipartite — there are no edges connecting performers to performers or movies to movies.

### Minimum spanning tree

Computing MST (minimum spanning tree) for undirected edge-weighted graphs.
Weight represents a cost (or time) that it takes to propagate through the edge
(airline map where weights are distances or fares).
Minimizing the cost is naturally of interest in such situations.

Worst-case order of growth for V vertices and E edges.

| algorithm | space | time
| ---       | ---   | ---
| [lazy Prim](https://godoc.org/github.com/marselester/alg/graph/mst#LazyPrim) | E | E log E
| eager Prim | V | E log V
| [Kruskal](https://godoc.org/github.com/marselester/alg/graph/mst#Kruskal) | E | E log E
| Fredman-Tarjan | V | E + V log V
| Chazelle | V | nearly E

## Digraph

Digraph-processing problems and solutions:

- **single and multiple-source reachability**
  ([depth-first search](https://godoc.org/github.com/marselester/alg/digraph#DepthFirstSearch))
  — find vertices that are reachable from a given vertex
- **single-source directed paths** (code is identical to undirected depth-first paths)
  — find a directed path from source to a target vertex
- **single-source shortest directed paths** (code is identical to undirected breadth-first paths)
  — find a shortest directed path from source to a target vertex
- **directed cycle detection**
  ([has cycle](https://godoc.org/github.com/marselester/alg/digraph#HasCycle))
  — in a scheduling problem certain jobs must be performed before certain others (precedence constraint)
  which is a topological sort problem. No solution exists if there is a directed cycle.
- **precedence-constrained scheduling**
  ([topological sort](https://godoc.org/github.com/marselester/alg/digraph#TopologicalSort))
- **strong connectivity**
  ([Kosaraju-Sharir algorithm](https://godoc.org/github.com/marselester/alg/digraph#StrongConnectedComponent))
  — two vertices v and w are strongly connected if there is a directed path from v to w, and from w to v.
  Strong connectivity helps to understand the structure of a digraph,
  highlighting interrelated sets of vertices (strong components).
  For example, you can model a food chain (predator-prey) to study ecological system.
  Another example is web content where a page represents a vertex, and edge represents a hyperlink.
- **all-pairs reachability**
  ([transitive closure](https://godoc.org/github.com/marselester/alg/digraph#TransitiveClosure))
  — is there a directed path from vertex v to w? Note, the pair of vertices v and w are not strongly connected.

### Shortest paths

Find a lowest-cost way to get from one vertex to another.
The classic Dijkstra's algorithm for the problems when weights are nonnegative:

- **single-source shortest paths**
  ([Dijkstra's algorithm](https://godoc.org/github.com/marselester/alg/digraph/spt#Dijkstra))
  — is there a directed path from s to a given target vertex t? If so, find a shortest path (total weight is minimal).
- **source-sink shortest paths** — find the shortest path from s to t.
  Use Dijkstra's algorithm, but terminate the search as soon as t comes off the priority queue.
- **all-pairs shortest paths**
  ([Dijkstra all-pairs](https://godoc.org/github.com/marselester/alg/digraph/spt#DijkstraAllPairs))
  — find a shortest path from s to t (all vertex pairs).

[Topological sort](https://godoc.org/github.com/marselester/alg/digraph/spt#Acyclic)
is a faster algorithm for acyclic edge-weighted digraphs that works even when edge weights can be negative.

The classic Bellman-Ford algorithm for use in the general case, when cycles may be present, edge weights may be negative,
and we need algorithms for finding negative-weight cycles and shortest paths in edge-weighted digraphs with no such cycles.

### Longest paths

Find a highest-cost way to get from one vertex to another (negative weights allowed).

- **single-source longest paths**
  [Topological sort](https://godoc.org/github.com/marselester/alg/digraph/lpt#Acyclic)
  — is there a directed path from s to a given target vertex t? If so, find a longest path (total weight is maximal).
- **parallel precedence-constrained scheduling**
  ([critical path method](https://godoc.org/github.com/marselester/alg/cmd/schedule))
  is equivalent to a longest-paths problem

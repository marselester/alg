# Algorithms

[![Documentation](https://godoc.org/github.com/marselester/alg?status.svg)](https://godoc.org/github.com/marselester/alg)
[![Go Report Card](https://goreportcard.com/badge/github.com/marselester/alg)](https://goreportcard.com/report/github.com/marselester/alg)

Go implementation of some examples from [Algorithms book](https://algs4.cs.princeton.edu).

Table of contents:

- [Sorting](#sorting)
  - [String sorts](#string-sorts)
- [Searching](#searching)
  - [String symbol-table](#string-symbol-table)
  - [Substring search](#substring-search)
- [Graphs](#graphs)

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

## Graphs

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

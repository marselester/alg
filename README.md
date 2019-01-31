# Algorithms

[![Documentation](https://godoc.org/github.com/marselester/alg?status.svg)](https://godoc.org/github.com/marselester/alg)
[![Go Report Card](https://goreportcard.com/badge/github.com/marselester/alg)](https://goreportcard.com/report/github.com/marselester/alg)

Go implementation of some examples from [Algorithms book](https://algs4.cs.princeton.edu).

## Sorting

| algorithm      | running time | notes
| ---            | ---          | ---
| selection sort | n²           |
| insertion sort | between n and n² (depends on order of items) | stable
| shellsort      | n log n ?    |
| quicksort      | n log n (probabilistic guarantee) | extra space lg n
| mergesort      | n log n      | stable, not in place, extra space n
| heapsort       | n log n      |

A sorting method is stable if it preserves the relative order of equal keys in the array.
For example, transactions sorted by location will still preserve order by timestamp.

It's reasonable to avoid the costs of using references and sort primitive types instead.

Quicksort is the fastest general-purpose sort:

- it has only a few instructions in its inner loop
- it does well with cache memories because it most often references data sequentially

Insertion sort is an excellent method for partially sorted arrays and is also
a fine method for tiny arrays (~15 items). These properties can be leveraged in
intermediate stages of mergesort and quicksort.

Heapsort is popular when space is tight (embedded systems), but it has poor cache performance:
array entries are rarely compared with nearby array entries, so the number of cache misses is far
higher than for quicksort, mergesort, shellsort.

## Searching

| data structure                  | pros | cons
| ---                             | ---  | ---
| linked list (sequential search) | best for tiny symbol tables | slow for large symbol tables
| ordered array (binary search)   | optimal search and space, order-based ops | slow insert
| binary search tree              | easy to implement, order-based ops | no guarantess, space for links
| balanced binary search tree     | optimal search and insert, order-based ops | space for links
| hash table                      | fast search/insert | hash for each type, space for links/empty

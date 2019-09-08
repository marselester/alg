/*
Package strsearch takes advantage of properties of strings to provide search methods
(symbol-table implementations) that can be more efficient than the general-purpose search methods.
They have the following performance characteristics even for huge tables:

	search hits take time proportional to the length of the search key
	searh misses involve examining only a few characters

*/
package strsearch

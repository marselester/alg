// Package qunion solves dynamic connectivity problem using quick-union algorithm
// that focuses on speeding up the union operation from quick-find solution.
// Best-case input for the client is linear running time; worst-case is quadratic.
package qunion

// Network represents dynamic connections in a network.
type Network struct {
	// sites is an array of sites (e.g., a computer in a network);
	// it's a parent-link representation of a forest of trees.
	// An index represents a site number, value is a "link" to another site which
	// belongs to the same component (root of a tree).
	sites []int
	// count is a number of components in the network.
	count int
}

// New creates a Network of size n. Its components have IDs that correspond
// to array index.
func New(n int) *Network {
	net := Network{
		sites: make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		net.sites[i] = i
	}
	return &net
}

// Find returns the component identifier for a given site p.
// It follows site links until reaching a root site (component ID) that
// has a link to itself.
func (net *Network) Find(p int) int {
	for {
		p = net.sites[p]
		if net.sites[p] == p {
			return p
		}
	}
}

// Connect adds a connection between p and q by merging components
// if the two sites are in different components. Each merge decrements
// the number of components by one.
//
// To combine the two components into one, find their roots and link
// one of the roots to another.
func (net *Network) Connect(p, q int) {
	pID := net.Find(p)
	qID := net.Find(q)
	if pID == qID {
		return
	}

	net.sites[pID] = qID

	net.count--
}

// IsConnected tells whether p and q are in the same component (they have the same root).
func (net *Network) IsConnected(p, q int) bool {
	return net.Find(p) == net.Find(q)
}

// Count returns number of components. Initially, there are n components,
// with each site in its own component.
func (net *Network) Count() int {
	return net.count
}

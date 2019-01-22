// Package wqunion solves dynamic connectivity problem by weighted quick-union algorithm.
// Rather than arbitrary connecting the second tree to the first in quick-union,
// we always connect smaller tree to the larger. For that we have to keep track of
// the size of each tree.
// A running time of a client program is logarithmic.
package wqunion

// Network represents dynamic connections in a network.
type Network struct {
	// sites is an array of sites (e.g., a computer in a network);
	// it's a parent-link representation of a forest of trees.
	// An index represents a site number, value is a "link" to another site which
	// belongs to the same component (root of a tree).
	sites []int
	// sizes represents the size of each tree (how many sites belong to component).
	sizes []int
	// count is a number of components in the network.
	count int
}

// New creates a Network of size n. Its components have IDs that correspond
// to array index. Initially each component has one site (size is one).
func New(n int) *Network {
	net := Network{
		sites: make([]int, n),
		sizes: make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		net.sites[i] = i
		net.sizes[i] = 1
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
// smaller tree to the larger.
func (net *Network) Connect(p, q int) {
	pID := net.Find(p)
	qID := net.Find(q)
	if pID == qID {
		return
	}

	if net.sizes[pID] < net.sizes[qID] {
		net.sites[pID] = qID
		net.sizes[qID] += net.sizes[pID]
	} else {
		net.sites[qID] = pID
		net.sizes[pID] += net.sizes[qID]
	}

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

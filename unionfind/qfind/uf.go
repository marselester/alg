// Package qfind solves dynamic connectivity problem using quick-find algorithm.
// For example, you can think of the integers p and q as belonging to mathematical sets.
// When we process a pair p q, we are asking whether they belong to the same set.
// If not, we unite p's set and q's set.
//
// The integers might represent computers in a large network,
// and the pairs -- connections in the network. Determine whether we need to establish
// a new direct connection for p and q to be able to communicate or
// whether we can use existing connections to set up a communication path.
// A set of connected computers form a component.
//
// A client program that constructs a network will have quadratic running time.
package qfind

// Network represents dynamic connections in a network.
type Network struct {
	// sites is an array of sites (e.g., a computer in a network).
	// An index represents a site number, value corresponds its component ID.
	// In other words, where a site belongs to.
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
// The ID is always one of the sites in the component.
func (net *Network) Find(p int) int {
	return net.sites[p]
}

// Connect adds a connection between p and q by merging components
// if the two sites are in different components. Each merge decrements
// the number of components by one.
//
// To combine the two components into one, we have to make all of the
// sites entries corresponding to both sets of sites the same ID.
func (net *Network) Connect(p, q int) {
	pID := net.Find(p)
	qID := net.Find(q)
	if pID == qID {
		return
	}

	for i := 0; i < len(net.sites); i++ {
		if net.sites[i] == pID {
			net.sites[i] = qID
		}
	}

	net.count--
}

// IsConnected tells whether p and q are in the same component.
// All sites in a component must have the same component ID.
func (net *Network) IsConnected(p, q int) bool {
	return net.Find(p) == net.Find(q)
}

// Count returns number of components. Initially, there are n components,
// with each site in its own component.
func (net *Network) Count() int {
	return net.count
}

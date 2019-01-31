package symboltable

// SequentialSearch represents a symbol table based on a linked list that contains
// unique keys and associated values. The average number of compares for a random search hit is n/2,
// worst case cost is n.
//
// Move-to-front heuristic can be used to make frequently accessed keys likely to be found early:
// on every search hit, move the key-value pair to the beginning of the list.
// Then move all pairs between the beginning and the vacated position to the right one position.
//
// The location of the most recently accessed key can be cached to optimize Get.
type SequentialSearch struct {
	first *node
}
type node struct {
	key   string
	value int
	next  *node
}

// Get scans through the list and compares the search key with key in each node.
func (st *SequentialSearch) Get(key string) int {
	for n := st.first; n != nil; n = n.next {
		if n.key == key {
			return n.value
		}
	}
	return -1
}

// Put scans through the list and updates the value associated with the search key.
// If key is not found, a new node is inserted at the beginning of the list.
func (st *SequentialSearch) Put(key string, value int) {
	for n := st.first; n != nil; n = n.next {
		if n.key == key {
			n.value = value
			return
		}
	}

	st.first = &node{
		key:   key,
		value: value,
		next:  st.first,
	}
}

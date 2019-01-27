package symboltable

// SequentialSearch represents a symbol table based on a linked list that contains
// unique keys and associated values. The average number of compares for a random search hit is n/2.
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

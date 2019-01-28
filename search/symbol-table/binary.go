package symboltable

// BinarySearch represents a symbol table based on a binary search algorithm.
// Keys are stored in an ordered array to reduce number of compares required for each search.
type BinarySearch struct {
	keys   []string
	values []int
}

// rank relies on binary search to find a key in the sorted key array.
// If a key is found, its array index is returned (index equals to number of keys smaller than the search key).
// Otherwise, it also returns the number of keys that are smaller than the search key.
func (st *BinarySearch) rank(key string) int {
	var lo, mid, hi int
	hi = len(st.keys) - 1

	for hi >= lo {
		mid = lo + (hi-lo)/2
		switch {
		case key == st.keys[mid]:
			return mid
		case key > st.keys[mid]:
			lo = mid + 1
		case key < st.keys[mid]:
			hi = mid - 1
		}
	}

	// lo always equals to number of keys that are smaller than the search key.
	return lo
}

// Get uses rank that tells at what index the key to be found.
// If the located key doesn't match, then it's not in the symbol table.
func (st *BinarySearch) Get(key string) int {
	i := st.rank(key)
	if i < len(st.keys) && st.keys[i] == key {
		return st.values[i]
	}
	return -1
}

// Put uses rank that tells at what index update the value and
// where to put the key when the key is not in the symbol table.
func (st *BinarySearch) Put(key string, value int) {
	if st.keys == nil {
		st.keys = append(st.keys, key)
		st.values = append(st.values, value)
	}

	i := st.rank(key)
	// Update the value when key is found.
	if i < len(st.keys) && st.keys[i] == key {
		st.values[i] = value
		return
	}

	// Move larger keys one position to the right (working from back to front) to make room
	// and insert the given key/value.
	// But firstly, grow the underlying arrays.
	hi := len(st.keys) - 1
	st.keys = append(st.keys, st.keys[hi])
	st.values = append(st.values, st.values[hi])
	for j := hi; j > i; j-- {
		st.keys[j] = st.keys[j-1]
		st.values[j] = st.values[j-1]
	}

	st.keys[i] = key
	st.values[i] = value
}
